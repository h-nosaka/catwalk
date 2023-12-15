package money_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

/*
// 残念ながら言語も合わせないとうまく取り出せない
a = 1234.5678
currencies = ["AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EEK","EGP","ERN","ETB","EUR","FJD","FKP","GBP","GEL","GGP","GHC","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RUR","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SKK","SLL","SOS","SRD","SSP","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRL","TRY","TTD","TWD","TZS","UAH","UGX","USD","UYU","UZS","VEF","VND","VUV","WST","XAF","XAG","XAU","XCD","XDR","XOF","XPF","YER","ZAR","ZMW","ZWD","ZWL"]
JSON.stringify(currencies.map(l => {
  return [l, a.toLocaleString("en", { style: 'currency', currency: l })]
}))
*/

const Currencies = `[["AED","AED 1,234.57"],["AFN","AFN 1,235"],["ALL","ALL 1,235"],["AMD","AMD 1,234.57"],["ANG","ANG 1,234.57"],["AOA","AOA 1,234.57"],["ARS","ARS 1,234.57"],["AUD","A$1,234.57"],["AWG","AWG 1,234.57"],["AZN","AZN 1,234.57"],["BAM","BAM 1,234.57"],["BBD","BBD 1,234.57"],["BDT","BDT 1,234.57"],["BGN","BGN 1,234.57"],["BHD","BHD 1,234.568"],["BIF","BIF 1,235"],["BMD","BMD 1,234.57"],["BND","BND 1,234.57"],["BOB","BOB 1,234.57"],["BRL","R$1,234.57"],["BSD","BSD 1,234.57"],["BTN","BTN 1,234.57"],["BWP","BWP 1,234.57"],["BYN","BYN 1,234.57"],["BYR","BYR 1,235"],["BZD","BZD 1,234.57"],["CAD","CA$1,234.57"],["CDF","CDF 1,234.57"],["CHF","CHF 1,234.57"],["CLF","CLF 1,234.5678"],["CLP","CLP 1,235"],["CNY","CN¥1,234.57"],["COP","COP 1,234.57"],["CRC","CRC 1,234.57"],["CUC","CUC 1,234.57"],["CUP","CUP 1,234.57"],["CVE","CVE 1,234.57"],["CZK","CZK 1,234.57"],["DJF","DJF 1,235"],["DKK","DKK 1,234.57"],["DOP","DOP 1,234.57"],["DZD","DZD 1,234.57"],["EEK","EEK 1,234.57"],["EGP","EGP 1,234.57"],["ERN","ERN 1,234.57"],["ETB","ETB 1,234.57"],["EUR","€1,234.57"],["FJD","FJD 1,234.57"],["FKP","FKP 1,234.57"],["GBP","£1,234.57"],["GEL","GEL 1,234.57"],["GGP","GGP 1,234.57"],["GHC","GHC 1,234.57"],["GHS","GHS 1,234.57"],["GIP","GIP 1,234.57"],["GMD","GMD 1,234.57"],["GNF","GNF 1,235"],["GTQ","GTQ 1,234.57"],["GYD","GYD 1,234.57"],["HKD","HK$1,234.57"],["HNL","HNL 1,234.57"],["HRK","HRK 1,234.57"],["HTG","HTG 1,234.57"],["HUF","HUF 1,234.57"],["IDR","IDR 1,234.57"],["ILS","₪1,234.57"],["IMP","IMP 1,234.57"],["INR","₹1,234.57"],["IQD","IQD 1,235"],["IRR","IRR 1,235"],["ISK","ISK 1,235"],["JEP","JEP 1,234.57"],["JMD","JMD 1,234.57"],["JOD","JOD 1,234.568"],["JPY","¥1,235"],["KES","KES 1,234.57"],["KGS","KGS 1,234.57"],["KHR","KHR 1,234.57"],["KMF","KMF 1,235"],["KPW","KPW 1,235"],["KRW","₩1,235"],["KWD","KWD 1,234.568"],["KYD","KYD 1,234.57"],["KZT","KZT 1,234.57"],["LAK","LAK 1,235"],["LBP","LBP 1,235"],["LKR","LKR 1,234.57"],["LRD","LRD 1,234.57"],["LSL","LSL 1,234.57"],["LTL","LTL 1,234.57"],["LVL","LVL 1,234.57"],["LYD","LYD 1,234.568"],["MAD","MAD 1,234.57"],["MDL","MDL 1,234.57"],["MGA","MGA 1,235"],["MKD","MKD 1,234.57"],["MMK","MMK 1,235"],["MNT","MNT 1,234.57"],["MOP","MOP 1,234.57"],["MUR","MUR 1,234.57"],["MVR","MVR 1,234.57"],["MWK","MWK 1,234.57"],["MXN","MX$1,234.57"],["MYR","MYR 1,234.57"],["MZN","MZN 1,234.57"],["NAD","NAD 1,234.57"],["NGN","NGN 1,234.57"],["NIO","NIO 1,234.57"],["NOK","NOK 1,234.57"],["NPR","NPR 1,234.57"],["NZD","NZ$1,234.57"],["OMR","OMR 1,234.568"],["PAB","PAB 1,234.57"],["PEN","PEN 1,234.57"],["PGK","PGK 1,234.57"],["PHP","₱1,234.57"],["PKR","PKR 1,234.57"],["PLN","PLN 1,234.57"],["PYG","PYG 1,235"],["QAR","QAR 1,234.57"],["RON","RON 1,234.57"],["RSD","RSD 1,235"],["RUB","RUB 1,234.57"],["RUR","RUR 1,234.57"],["RWF","RWF 1,235"],["SAR","SAR 1,234.57"],["SBD","SBD 1,234.57"],["SCR","SCR 1,234.57"],["SDG","SDG 1,234.57"],["SEK","SEK 1,234.57"],["SGD","SGD 1,234.57"],["SHP","SHP 1,234.57"],["SKK","SKK 1,234.57"],["SLL","SLL 1,235"],["SOS","SOS 1,235"],["SRD","SRD 1,234.57"],["SSP","SSP 1,234.57"],["STD","STD 1,235"],["SVC","SVC 1,234.57"],["SYP","SYP 1,235"],["SZL","SZL 1,234.57"],["THB","THB 1,234.57"],["TJS","TJS 1,234.57"],["TMT","TMT 1,234.57"],["TND","TND 1,234.568"],["TOP","TOP 1,234.57"],["TRL","TRL 1,235"],["TRY","TRY 1,234.57"],["TTD","TTD 1,234.57"],["TWD","NT$1,234.57"],["TZS","TZS 1,234.57"],["UAH","UAH 1,234.57"],["UGX","UGX 1,235"],["USD","$1,234.57"],["UYU","UYU 1,234.57"],["UZS","UZS 1,234.57"],["VEF","VEF 1,234.57"],["VND","₫1,235"],["VUV","VUV 1,235"],["WST","WST 1,234.57"],["XAF","FCFA 1,235"],["XAG","XAG 1,234.57"],["XAU","XAU 1,234.57"],["XCD","EC$1,234.57"],["XDR","XDR 1,234.57"],["XOF","F CFA 1,235"],["XPF","CFPF 1,235"],["YER","YER 1,235"],["ZAR","ZAR 1,234.57"],["ZMW","ZMW 1,234.57"],["ZWD","ZWD 1,235"],["ZWL","ZWL 1,234.57"]]`

func TestMakeCurrencies(t *testing.T) {
	data := Dump{}
	if err := json.Unmarshal([]byte(Currencies), &data); err != nil {
		t.Error(err)
	}
	for _, item := range data {
		key := item[0]
		currency := strings.Split(item[1], " ")
		if len(currency) == 1 {
			currency = strings.Split(item[1], "1,2")
		}
		comma := strings.Split(item[1], ".")
		fraction := 0
		if len(comma) > 1 {
			fraction = len(comma[1])
		}
		fmt.Printf(`%s: {Code: %s, Fraction: %d, Grapheme: "%s"}, // %s`, key, key, fraction, currency[0], item[1])
		fmt.Print("\n")
	}
	t.Error("test")
}
