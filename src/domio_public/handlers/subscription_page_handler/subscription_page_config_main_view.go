package subscription_page_handler

func GetSubscriptionConfigMainViewTemplate() string {
    return `{{ define "subscription_config_mainview_template"}}

                {{with .PageData}}
                    <input class="subscription-id-input" type="hidden" value="{{.Subscription.Id}}"></input>

                    <div class="b-subscription-config-mainview">

                        <div class="b-records-list-header">
                            <div class="header-name">Name</div>
                            <div class="header-type">Type</div>
                            <div class="header-value">Value</div>
                            <div class="header-eth">Evaluate Target Health</div>
                            <div class="header-hcid">Health Check ID</div>
                            <div class="header-ttl">TTL</div>
                            <div class="header-region">Region</div>
                            <div class="header-weight">Weight</div>
                            <div class="header-geolocation">Geolocation</div>
                            <div class="header-setid">Set ID</div>
                        </div>


                        <div class="b-records-list-container">

                            {{range .Records}}

                                <div class="b-record-row">

                                    <div class="cell cell-name">{{.Name}}</div>
                                    <div class="cell cell-type">{{.Type}}</div>

                                    <div class="cell cell-value" data-type="{{.Type}}">

                                        {{range .ResourceRecords}}

                                            <div class="b-record-value-row">{{.Value}}</div>

                                        {{end}}

                                    </div>

                                    <div class="cell cell-eth">!eth!</div>

                                    <div class="cell cell-hcid">{{.HealthCheckId}}</div>
                                    <div class="cell cell-ttl">{{.TTL}}</div>
                                    <div class="cell cell-region">{{.Region}}</div>
                                    <div class="cell cell-weight">{{.Weight}}</div>
                                    <div class="cell cell-geolocation">{{if .GeoLocation}}{{.GeoLocation.ContinentCode}}{{end}}</div>
                                    <div class="cell cell-setid">{{.SetIdentifier}}</div>

                                </div>

                            {{end}}

                        </div>

                        <form class="b-subscription-record-form">
                            <select class="value-type-select">
                                <option value="a">A – IPv4 address</option>
                                <option value="cname">CNAME – Canonical name</option>
                                <option value="mx">MX – Mail exchange</option>
                                <option value="aaaa">AAAA – IPv6 address</option>
                                <option value="txt">TXT – Text</option>
                                <option value="ptr">PTR – Pointer</option>
                                <option value="srv">SRV – Service locator</option>
                                <option value="spf">SPF – Sender Policy Framework</option>
                                <option value="naptr">NAPTR – Name Authority Pointer</option>
                                <option value="ns">NS – Name server</option>
                                <option value="soa" disabled="disabled" title="Only one SOA record can exist in a zone, you can select it to edit its properties.">SOA – Start of authority</option>
                            </select>

                            <input class="value-input"></input>
                            <input type="submit" class="update-subscription-button" type="submit" value="Add"></input>
                        </form>

                    </div>

                {{end}}

            {{end}}
        `
}