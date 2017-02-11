package add_payment_source_page_handler

import (
    "html/template"
    "log"
)

func GetAddPaymentSourcePageTemplate(parsedTemplate *template.Template) {
    _, err := parsedTemplate.New("main_template").Parse(`{{define "main_template"}}
                                                            <form>
                                                                {{with .PageData.DomainInfo}}
                                                                    <div data-domain="{{.Name}}">
                                                                        <h2>{{.Name}}</h2>

                                                                        <div>
                                                                            <label>Name on the card:</label>
                                                                            <input name="name_on_the_card" type="text" value="John Doe"></input>
                                                                        </div>

                                                                        <div>
                                                                            <label>Expiry Month:</label>
                                                                            <input name="expiry_month" type="number" value="12"></input>
                                                                        </div>

                                                                        <div>
                                                                            <label>Expiry Year:</label>
                                                                            <input name="expiry_year" type="number" value="2018"></input>
                                                                        </div>

                                                                        <div>
                                                                            <label>CVC:</label>
                                                                            <input name="cvc" type="number" value="123"></input>
                                                                        </div>

                                                                        <div>
                                                                            <label>Card Number:</label>
                                                                            <input name="card_number" type="number" value="4242424242424242"></input>
                                                                        </div>

                                                                        <button class="b-add-payment-source-button">Add card</button>

                                                                    </div>
                                                                {{end}}
                                                            </form>
                                                         {{end}}`)
    if (err != nil) {
        log.Print(err)
    }
}