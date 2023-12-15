package money

type Currency string

const (
	AED Currency = "AED"
	AFN Currency = "AFN"
	ALL Currency = "ALL"
	AMD Currency = "AMD"
	ANG Currency = "ANG"
	AOA Currency = "AOA"
	ARS Currency = "ARS"
	AUD Currency = "AUD"
	AWG Currency = "AWG"
	AZN Currency = "AZN"
	BAM Currency = "BAM"
	BBD Currency = "BBD"
	BDT Currency = "BDT"
	BGN Currency = "BGN"
	BHD Currency = "BHD"
	BIF Currency = "BIF"
	BMD Currency = "BMD"
	BND Currency = "BND"
	BOB Currency = "BOB"
	BRL Currency = "BRL"
	BSD Currency = "BSD"
	BTN Currency = "BTN"
	BWP Currency = "BWP"
	BYN Currency = "BYN"
	BYR Currency = "BYR"
	BZD Currency = "BZD"
	CAD Currency = "CAD"
	CDF Currency = "CDF"
	CHF Currency = "CHF"
	CLF Currency = "CLF"
	CLP Currency = "CLP"
	CNY Currency = "CNY"
	COP Currency = "COP"
	CRC Currency = "CRC"
	CUC Currency = "CUC"
	CUP Currency = "CUP"
	CVE Currency = "CVE"
	CZK Currency = "CZK"
	DJF Currency = "DJF"
	DKK Currency = "DKK"
	DOP Currency = "DOP"
	DZD Currency = "DZD"
	EEK Currency = "EEK"
	EGP Currency = "EGP"
	ERN Currency = "ERN"
	ETB Currency = "ETB"
	EUR Currency = "EUR"
	FJD Currency = "FJD"
	FKP Currency = "FKP"
	GBP Currency = "GBP"
	GEL Currency = "GEL"
	GGP Currency = "GGP"
	GHC Currency = "GHC"
	GHS Currency = "GHS"
	GIP Currency = "GIP"
	GMD Currency = "GMD"
	GNF Currency = "GNF"
	GTQ Currency = "GTQ"
	GYD Currency = "GYD"
	HKD Currency = "HKD"
	HNL Currency = "HNL"
	HRK Currency = "HRK"
	HTG Currency = "HTG"
	HUF Currency = "HUF"
	IDR Currency = "IDR"
	ILS Currency = "ILS"
	IMP Currency = "IMP"
	INR Currency = "INR"
	IQD Currency = "IQD"
	IRR Currency = "IRR"
	ISK Currency = "ISK"
	JEP Currency = "JEP"
	JMD Currency = "JMD"
	JOD Currency = "JOD"
	JPY Currency = "JPY"
	KES Currency = "KES"
	KGS Currency = "KGS"
	KHR Currency = "KHR"
	KMF Currency = "KMF"
	KPW Currency = "KPW"
	KRW Currency = "KRW"
	KWD Currency = "KWD"
	KYD Currency = "KYD"
	KZT Currency = "KZT"
	LAK Currency = "LAK"
	LBP Currency = "LBP"
	LKR Currency = "LKR"
	LRD Currency = "LRD"
	LSL Currency = "LSL"
	LTL Currency = "LTL"
	LVL Currency = "LVL"
	LYD Currency = "LYD"
	MAD Currency = "MAD"
	MDL Currency = "MDL"
	MGA Currency = "MGA"
	MKD Currency = "MKD"
	MMK Currency = "MMK"
	MNT Currency = "MNT"
	MOP Currency = "MOP"
	MUR Currency = "MUR"
	MVR Currency = "MVR"
	MWK Currency = "MWK"
	MXN Currency = "MXN"
	MYR Currency = "MYR"
	MZN Currency = "MZN"
	NAD Currency = "NAD"
	NGN Currency = "NGN"
	NIO Currency = "NIO"
	NOK Currency = "NOK"
	NPR Currency = "NPR"
	NZD Currency = "NZD"
	OMR Currency = "OMR"
	PAB Currency = "PAB"
	PEN Currency = "PEN"
	PGK Currency = "PGK"
	PHP Currency = "PHP"
	PKR Currency = "PKR"
	PLN Currency = "PLN"
	PYG Currency = "PYG"
	QAR Currency = "QAR"
	RON Currency = "RON"
	RSD Currency = "RSD"
	RUB Currency = "RUB"
	RUR Currency = "RUR"
	RWF Currency = "RWF"
	SAR Currency = "SAR"
	SBD Currency = "SBD"
	SCR Currency = "SCR"
	SDG Currency = "SDG"
	SEK Currency = "SEK"
	SGD Currency = "SGD"
	SHP Currency = "SHP"
	SKK Currency = "SKK"
	SLL Currency = "SLL"
	SOS Currency = "SOS"
	SRD Currency = "SRD"
	SSP Currency = "SSP"
	STD Currency = "STD"
	SVC Currency = "SVC"
	SYP Currency = "SYP"
	SZL Currency = "SZL"
	THB Currency = "THB"
	TJS Currency = "TJS"
	TMT Currency = "TMT"
	TND Currency = "TND"
	TOP Currency = "TOP"
	TRL Currency = "TRL"
	TRY Currency = "TRY"
	TTD Currency = "TTD"
	TWD Currency = "TWD"
	TZS Currency = "TZS"
	UAH Currency = "UAH"
	UGX Currency = "UGX"
	USD Currency = "USD"
	UYU Currency = "UYU"
	UZS Currency = "UZS"
	VEF Currency = "VEF"
	VND Currency = "VND"
	VUV Currency = "VUV"
	WST Currency = "WST"
	XAF Currency = "XAF"
	XAG Currency = "XAG"
	XAU Currency = "XAU"
	XCD Currency = "XCD"
	XDR Currency = "XDR"
	XOF Currency = "XOF"
	XPF Currency = "XPF"
	YER Currency = "YER"
	ZAR Currency = "ZAR"
	ZMW Currency = "ZMW"
	ZWD Currency = "ZWD"
	ZWL Currency = "ZWL"
)

func (p Currency) String() string {
	return string(p)
}

type CurrencyData struct {
	Code     Currency
	Fraction int
	Grapheme string
}
type Currencies map[Currency]*CurrencyData

var currencies = Currencies{
	AED: {Code: AED, Fraction: 2, Grapheme: ".\u062f.\u0625"},
	AFN: {Code: AFN, Fraction: 2, Grapheme: "\u060b"},
	ALL: {Code: ALL, Fraction: 2, Grapheme: "L"},
	AMD: {Code: AMD, Fraction: 2, Grapheme: "\u0564\u0580."},
	ANG: {Code: ANG, Fraction: 2, Grapheme: "\u0192"},
	AOA: {Code: AOA, Fraction: 2, Grapheme: "Kz"},
	ARS: {Code: ARS, Fraction: 2, Grapheme: "$"},
	AUD: {Code: AUD, Fraction: 2, Grapheme: "$"},
	AWG: {Code: AWG, Fraction: 2, Grapheme: "\u0192"},
	AZN: {Code: AZN, Fraction: 2, Grapheme: "\u20bc"},
	BAM: {Code: BAM, Fraction: 2, Grapheme: "KM"},
	BBD: {Code: BBD, Fraction: 2, Grapheme: "$"},
	BDT: {Code: BDT, Fraction: 2, Grapheme: "\u09f3"},
	BGN: {Code: BGN, Fraction: 2, Grapheme: "\u043b\u0432"},
	BHD: {Code: BHD, Fraction: 3, Grapheme: ".\u062f.\u0628"},
	BIF: {Code: BIF, Fraction: 0, Grapheme: "Fr"},
	BMD: {Code: BMD, Fraction: 2, Grapheme: "$"},
	BND: {Code: BND, Fraction: 2, Grapheme: "$"},
	BOB: {Code: BOB, Fraction: 2, Grapheme: "Bs."},
	BRL: {Code: BRL, Fraction: 2, Grapheme: "R$"},
	BSD: {Code: BSD, Fraction: 2, Grapheme: "$"},
	BTN: {Code: BTN, Fraction: 2, Grapheme: "Nu."},
	BWP: {Code: BWP, Fraction: 2, Grapheme: "P"},
	BYN: {Code: BYN, Fraction: 2, Grapheme: "p."},
	BYR: {Code: BYR, Fraction: 0, Grapheme: "p."},
	BZD: {Code: BZD, Fraction: 2, Grapheme: "BZ$"},
	CAD: {Code: CAD, Fraction: 2, Grapheme: "$"},
	CDF: {Code: CDF, Fraction: 2, Grapheme: "FC"},
	CHF: {Code: CHF, Fraction: 2, Grapheme: "CHF"},
	CLF: {Code: CLF, Fraction: 4, Grapheme: "UF"},
	CLP: {Code: CLP, Fraction: 0, Grapheme: "$"},
	CNY: {Code: CNY, Fraction: 2, Grapheme: "\u5143"},
	COP: {Code: COP, Fraction: 2, Grapheme: "$"},
	CRC: {Code: CRC, Fraction: 2, Grapheme: "\u20a1"},
	CUC: {Code: CUC, Fraction: 2, Grapheme: "$"},
	CUP: {Code: CUP, Fraction: 2, Grapheme: "$MN"},
	CVE: {Code: CVE, Fraction: 2, Grapheme: "$"},
	CZK: {Code: CZK, Fraction: 2, Grapheme: "K\u010d"},
	DJF: {Code: DJF, Fraction: 0, Grapheme: "Fdj"},
	DKK: {Code: DKK, Fraction: 2, Grapheme: "kr"},
	DOP: {Code: DOP, Fraction: 2, Grapheme: "RD$"},
	DZD: {Code: DZD, Fraction: 2, Grapheme: ".\u062f.\u062c"},
	EEK: {Code: EEK, Fraction: 2, Grapheme: "kr"},
	EGP: {Code: EGP, Fraction: 2, Grapheme: "\u00a3"},
	ERN: {Code: ERN, Fraction: 2, Grapheme: "Nfk"},
	ETB: {Code: ETB, Fraction: 2, Grapheme: "Br"},
	EUR: {Code: EUR, Fraction: 2, Grapheme: "\u20ac"},
	FJD: {Code: FJD, Fraction: 2, Grapheme: "$"},
	FKP: {Code: FKP, Fraction: 2, Grapheme: "\u00a3"},
	GBP: {Code: GBP, Fraction: 2, Grapheme: "\u00a3"},
	GEL: {Code: GEL, Fraction: 2, Grapheme: "\u10da"},
	GGP: {Code: GGP, Fraction: 2, Grapheme: "\u00a3"},
	GHC: {Code: GHC, Fraction: 2, Grapheme: "\u00a2"},
	GHS: {Code: GHS, Fraction: 2, Grapheme: "\u20b5"},
	GIP: {Code: GIP, Fraction: 2, Grapheme: "\u00a3"},
	GMD: {Code: GMD, Fraction: 2, Grapheme: "D"},
	GNF: {Code: GNF, Fraction: 0, Grapheme: "FG"},
	GTQ: {Code: GTQ, Fraction: 2, Grapheme: "Q"},
	GYD: {Code: GYD, Fraction: 2, Grapheme: "$"},
	HKD: {Code: HKD, Fraction: 2, Grapheme: "$"},
	HNL: {Code: HNL, Fraction: 2, Grapheme: "L"},
	HRK: {Code: HRK, Fraction: 2, Grapheme: "kn"},
	HTG: {Code: HTG, Fraction: 2, Grapheme: "G"},
	HUF: {Code: HUF, Fraction: 2, Grapheme: "Ft"},
	IDR: {Code: IDR, Fraction: 2, Grapheme: "Rp"},
	ILS: {Code: ILS, Fraction: 2, Grapheme: "\u20aa"},
	IMP: {Code: IMP, Fraction: 2, Grapheme: "\u00a3"},
	INR: {Code: INR, Fraction: 2, Grapheme: "\u20b9"},
	IQD: {Code: IQD, Fraction: 3, Grapheme: ".\u062f.\u0639"},
	IRR: {Code: IRR, Fraction: 2, Grapheme: "\ufdfc"},
	ISK: {Code: ISK, Fraction: 0, Grapheme: "kr"},
	JEP: {Code: JEP, Fraction: 2, Grapheme: "\u00a3"},
	JMD: {Code: JMD, Fraction: 2, Grapheme: "J$"},
	JOD: {Code: JOD, Fraction: 3, Grapheme: ".\u062f.\u0625"},
	JPY: {Code: JPY, Fraction: 0, Grapheme: "\u00a5"},
	KES: {Code: KES, Fraction: 2, Grapheme: "KSh"},
	KGS: {Code: KGS, Fraction: 2, Grapheme: "\u0441\u043e\u043c"},
	KHR: {Code: KHR, Fraction: 2, Grapheme: "\u17db"},
	KMF: {Code: KMF, Fraction: 0, Grapheme: "CF"},
	KPW: {Code: KPW, Fraction: 2, Grapheme: "\u20a9"},
	KRW: {Code: KRW, Fraction: 0, Grapheme: "\u20a9"},
	KWD: {Code: KWD, Fraction: 3, Grapheme: ".\u062f.\u0643"},
	KYD: {Code: KYD, Fraction: 2, Grapheme: "$"},
	KZT: {Code: KZT, Fraction: 2, Grapheme: "\u20b8"},
	LAK: {Code: LAK, Fraction: 2, Grapheme: "\u20ad"},
	LBP: {Code: LBP, Fraction: 2, Grapheme: "\u00a3"},
	LKR: {Code: LKR, Fraction: 2, Grapheme: "\u20a8"},
	LRD: {Code: LRD, Fraction: 2, Grapheme: "$"},
	LSL: {Code: LSL, Fraction: 2, Grapheme: "L"},
	LTL: {Code: LTL, Fraction: 2, Grapheme: "Lt"},
	LVL: {Code: LVL, Fraction: 2, Grapheme: "Ls"},
	LYD: {Code: LYD, Fraction: 3, Grapheme: ".\u062f.\u0644"},
	MAD: {Code: MAD, Fraction: 2, Grapheme: ".\u062f.\u0645"},
	MDL: {Code: MDL, Fraction: 2, Grapheme: "lei"},
	MGA: {Code: MGA, Fraction: 2, Grapheme: "Ar"},
	MKD: {Code: MKD, Fraction: 2, Grapheme: "\u0434\u0435\u043d"},
	MMK: {Code: MMK, Fraction: 2, Grapheme: "K"},
	MNT: {Code: MNT, Fraction: 2, Grapheme: "\u20ae"},
	MOP: {Code: MOP, Fraction: 2, Grapheme: "P"},
	MUR: {Code: MUR, Fraction: 2, Grapheme: "\u20a8"},
	MVR: {Code: MVR, Fraction: 2, Grapheme: "MVR"},
	MWK: {Code: MWK, Fraction: 2, Grapheme: "MK"},
	MXN: {Code: MXN, Fraction: 2, Grapheme: "$"},
	MYR: {Code: MYR, Fraction: 2, Grapheme: "RM"},
	MZN: {Code: MZN, Fraction: 2, Grapheme: "MT"},
	NAD: {Code: NAD, Fraction: 2, Grapheme: "$"},
	NGN: {Code: NGN, Fraction: 2, Grapheme: "\u20a6"},
	NIO: {Code: NIO, Fraction: 2, Grapheme: "C$"},
	NOK: {Code: NOK, Fraction: 2, Grapheme: "kr"},
	NPR: {Code: NPR, Fraction: 2, Grapheme: "\u20a8"},
	NZD: {Code: NZD, Fraction: 2, Grapheme: "$"},
	OMR: {Code: OMR, Fraction: 3, Grapheme: "\ufdfc"},
	PAB: {Code: PAB, Fraction: 2, Grapheme: "B/."},
	PEN: {Code: PEN, Fraction: 2, Grapheme: "S/"},
	PGK: {Code: PGK, Fraction: 2, Grapheme: "K"},
	PHP: {Code: PHP, Fraction: 2, Grapheme: "\u20b1"},
	PKR: {Code: PKR, Fraction: 2, Grapheme: "\u20a8"},
	PLN: {Code: PLN, Fraction: 2, Grapheme: "z\u0142"},
	PYG: {Code: PYG, Fraction: 0, Grapheme: "Gs"},
	QAR: {Code: QAR, Fraction: 2, Grapheme: "\ufdfc"},
	RON: {Code: RON, Fraction: 2, Grapheme: "lei"},
	RSD: {Code: RSD, Fraction: 2, Grapheme: "\u0414\u0438\u043d."},
	RUB: {Code: RUB, Fraction: 2, Grapheme: "\u20bd"},
	RUR: {Code: RUR, Fraction: 2, Grapheme: "\u20bd"},
	RWF: {Code: RWF, Fraction: 0, Grapheme: "FRw"},
	SAR: {Code: SAR, Fraction: 2, Grapheme: "\ufdfc"},
	SBD: {Code: SBD, Fraction: 2, Grapheme: "$"},
	SCR: {Code: SCR, Fraction: 2, Grapheme: "\u20a8"},
	SDG: {Code: SDG, Fraction: 2, Grapheme: "\u00a3"},
	SEK: {Code: SEK, Fraction: 2, Grapheme: "kr"},
	SGD: {Code: SGD, Fraction: 2, Grapheme: "$"},
	SHP: {Code: SHP, Fraction: 2, Grapheme: "\u00a3"},
	SKK: {Code: SKK, Fraction: 2, Grapheme: "Sk"},
	SLL: {Code: SLL, Fraction: 2, Grapheme: "Le"},
	SOS: {Code: SOS, Fraction: 2, Grapheme: "Sh"},
	SRD: {Code: SRD, Fraction: 2, Grapheme: "$"},
	SSP: {Code: SSP, Fraction: 2, Grapheme: "\u00a3"},
	STD: {Code: STD, Fraction: 2, Grapheme: "Db"},
	SVC: {Code: SVC, Fraction: 2, Grapheme: "\u20a1"},
	SYP: {Code: SYP, Fraction: 2, Grapheme: "\u00a3"},
	SZL: {Code: SZL, Fraction: 2, Grapheme: "\u00a3"},
	THB: {Code: THB, Fraction: 2, Grapheme: "\u0e3f"},
	TJS: {Code: TJS, Fraction: 2, Grapheme: "SM"},
	TMT: {Code: TMT, Fraction: 2, Grapheme: "T"},
	TND: {Code: TND, Fraction: 3, Grapheme: ".\u062f.\u062a"},
	TOP: {Code: TOP, Fraction: 2, Grapheme: "T$"},
	TRL: {Code: TRL, Fraction: 2, Grapheme: "\u20a4"},
	TRY: {Code: TRY, Fraction: 2, Grapheme: "\u20ba"},
	TTD: {Code: TTD, Fraction: 2, Grapheme: "TT$"},
	TWD: {Code: TWD, Fraction: 2, Grapheme: "NT$"},
	TZS: {Code: TZS, Fraction: 0, Grapheme: "TSh"},
	UAH: {Code: UAH, Fraction: 2, Grapheme: "\u20b4"},
	UGX: {Code: UGX, Fraction: 0, Grapheme: "USh"},
	USD: {Code: USD, Fraction: 2, Grapheme: "$"},
	UYU: {Code: UYU, Fraction: 2, Grapheme: "$U"},
	UZS: {Code: UZS, Fraction: 2, Grapheme: "so\u2019m"},
	VEF: {Code: VEF, Fraction: 2, Grapheme: "Bs"},
	VND: {Code: VND, Fraction: 0, Grapheme: "\u20ab"},
	VUV: {Code: VUV, Fraction: 0, Grapheme: "Vt"},
	WST: {Code: WST, Fraction: 2, Grapheme: "T"},
	XAF: {Code: XAF, Fraction: 0, Grapheme: "Fr"},
	XAG: {Code: XAG, Fraction: 0, Grapheme: "oz t"},
	XAU: {Code: XAU, Fraction: 0, Grapheme: "oz t"},
	XCD: {Code: XCD, Fraction: 2, Grapheme: "$"},
	XDR: {Code: XDR, Fraction: 0, Grapheme: "SDR"},
	XOF: {Code: XOF, Fraction: 0, Grapheme: "CFA"},
	XPF: {Code: XPF, Fraction: 0, Grapheme: "₣"},
	YER: {Code: YER, Fraction: 2, Grapheme: "\ufdfc"},
	ZAR: {Code: ZAR, Fraction: 2, Grapheme: "R"},
	ZMW: {Code: ZMW, Fraction: 2, Grapheme: "ZK"},
	ZWD: {Code: ZWD, Fraction: 2, Grapheme: "Z$"},
	ZWL: {Code: ZWL, Fraction: 2, Grapheme: "Z$"},
}

func (c Currencies) CurrencyByCode(code Currency) *CurrencyData {
	sc, ok := c[code]
	if !ok {
		return nil
	}

	return sc
}

func GetCurrency(code Currency) *CurrencyData {
	return currencies.CurrencyByCode(Currency(code))
}
