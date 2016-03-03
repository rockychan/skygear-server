package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	skyplugin "github.com/oursky/skygear/plugin"
	"github.com/oursky/skygear/plugin/common"
	"github.com/oursky/skygear/skyconfig"
	"github.com/oursky/skygear/skydb"
	"github.com/oursky/skygear/skydb/skyconv"
	"golang.org/x/net/context"
)

type httpTransport struct {
	URL         string
	Path        string
	Args        []string
	initHandler skyplugin.TransportInitHandler
	state       skyplugin.TransportState
	httpClient  http.Client
	config      skyconfig.Configuration
}

func (p *httpTransport) sendRequest(req *http.Request) (out []byte, err error) {
	httpresp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer httpresp.Body.Close()

	return ioutil.ReadAll(httpresp.Body)
}

func (p *httpTransport) execute(req *http.Request) (out []byte, err error) {
	data, err := p.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result json.RawMessage   `json:"result"`
		Err    *common.ExecError `json:"error"`
	}

	jsonErr := json.Unmarshal(data, &resp)
	if jsonErr != nil {
		err = fmt.Errorf("failed to parse response: %v", jsonErr)
		return
	}

	if resp.Err != nil {
		err = resp.Err
		return
	}

	out = resp.Result
	return
}

func (p *httpTransport) State() skyplugin.TransportState {
	return p.state
}

func (p *httpTransport) SetInitHandler(f skyplugin.TransportInitHandler) {
	p.initHandler = f
}

func (p *httpTransport) setState(state skyplugin.TransportState) {
	if state != p.state {
		oldState := p.state
		p.state = state
		log.Infof("Transport state changes from %v to %v.", oldState, p.state)
	}
}

func (p *httpTransport) RequestInit() {
	out, err := p.RunInit()
	if p.initHandler != nil {
		handlerError := p.initHandler(out, err)
		if err != nil || handlerError != nil {
			p.setState(skyplugin.TransportStateError)
			return
		}
	}
	p.setState(skyplugin.TransportStateReady)
}

func (p *httpTransport) RunInit() (out []byte, err error) {
	url := fmt.Sprintf("%s/init", p.Path)
	in, err := json.Marshal(struct {
		Config skyconfig.Configuration `json:"config"`
	}{p.config})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(in))
	if err != nil {
		return nil, err
	}
	return p.sendRequest(req)
}

func (p *httpTransport) RunLambda(ctx context.Context, name string, in []byte) (out []byte, err error) {
	url := fmt.Sprintf("%s/op/%s", p.Path, name)
	req, err := http.NewRequest("POST", url, bytes.NewReader(in))
	if err != nil {
		return nil, err
	}
	pluginCtx := skyplugin.ContextMap(ctx)
	encodedCtx, err := common.EncodeBase64JSON(pluginCtx)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Skygear-Plugin-Context", encodedCtx)
	return p.execute(req)
}

func (p *httpTransport) RunHandler(name string, in []byte) (out []byte, err error) {
	return nil, fmt.Errorf("not implemented")
}

func (p *httpTransport) RunHook(ctx context.Context, hookName string, record *skydb.Record, originalRecord *skydb.Record) (*skydb.Record, error) {
	param := map[string]interface{}{
		"record":   (*skyconv.JSONRecord)(record),
		"original": (*skyconv.JSONRecord)(originalRecord),
	}
	in, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal record: %v", err)
	}

	url := fmt.Sprintf("%s/hook/%s", p.Path, hookName)
	req, err := http.NewRequest("POST", url, bytes.NewReader(in))
	if err != nil {
		return nil, err
	}
	pluginCtx := skyplugin.ContextMap(ctx)
	encodedCtx, err := common.EncodeBase64JSON(pluginCtx)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Skygear-Plugin-Context", encodedCtx)

	out, err := p.execute(req)
	if err != nil {
		return nil, err
	}

	var recordout skydb.Record
	if err := json.Unmarshal(out, (*skyconv.JSONRecord)(&recordout)); err != nil {
		log.WithField("data", string(out)).Error("failed to unmarshal record")
		return nil, fmt.Errorf("failed to unmarshal record: %v", err)
	}
	recordout.OwnerID = record.OwnerID
	recordout.CreatedAt = record.CreatedAt
	recordout.CreatorID = record.CreatorID
	recordout.UpdatedAt = record.UpdatedAt
	recordout.UpdaterID = record.UpdaterID

	return &recordout, nil
}

func (p *httpTransport) RunTimer(name string, in []byte) (out []byte, err error) {
	url := fmt.Sprintf("%s/timer/%s", p.Path, name)
	req, err := http.NewRequest("POST", url, bytes.NewReader(in))
	if err != nil {
		return nil, err
	}
	return p.execute(req)
}

func (p *httpTransport) RunProvider(request *skyplugin.AuthRequest) (*skyplugin.AuthResponse, error) {
	req := map[string]interface{}{
		"auth_data": request.AuthData,
	}

	in, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal auth request: %v", err)
	}

	url := fmt.Sprintf("%s/provider/%s/%s", p.Path, request.ProviderName, request.Action)
	httpreq, err := http.NewRequest("POST", url, bytes.NewReader(in))
	if err != nil {
		return nil, err
	}
	out, err := p.execute(httpreq)

	resp := skyplugin.AuthResponse{}

	err = json.Unmarshal(out, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return &resp, nil
}

type httpTransportFactory struct {
}

func (f httpTransportFactory) Open(path string, args []string, config skyconfig.Configuration) (transport skyplugin.Transport) {
	transport = &httpTransport{
		Path:       path,
		Args:       args,
		state:      skyplugin.TransportStateUninitialized,
		httpClient: http.Client{},
		config:     config,
	}
	return
}

func init() {
	skyplugin.RegisterTransport("http", httpTransportFactory{})
}