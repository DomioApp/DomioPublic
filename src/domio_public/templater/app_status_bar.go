package templater

import (
    "html/template"
    "bytes"
    "domio_public/components/api"
    "log"
)

func GetAppStatusInfoBar(appStatusInfo api.AppStatusInfo) template.HTML {

    t := template.New("AppStatusInfoBar")

    output, _ := t.Parse(`
                            <div class="b-app-status-info">
                                <div>Version: <span>{{.Version}}</span></div>
                                <div>Built: <span>{{.BuildAgo}} ago</span></div>
                                <div>Buildstamp: <span>{{.Buildstamp}}</span></div>
                                <div>BuildDateTime: <span>{{.BuildDateTime}}</span></div>
                                <div>Hash: <span>{{.Hash}}</span></div>
                            </div>
                        `)

    var doc bytes.Buffer

    output.Execute(&doc, appStatusInfo)

    return template.HTML(doc.String())
}