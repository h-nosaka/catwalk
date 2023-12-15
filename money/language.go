package money

type Language string

const (
	IsIS      Language = "is-IS"
	Is        Language = "is"
	GaIE      Language = "ga-IE"
	Ga        Language = "ga"
	AzCyrl    Language = "az-Cyrl"
	AzCyrlAZ  Language = "az-Cyrl-AZ"
	AzLatn    Language = "az-Latn"
	AzLatnAZ  Language = "az-Latn-AZ"
	Az        Language = "az"
	AsIN      Language = "as-IN"
	As        Language = "as"
	AfZA      Language = "af-ZA"
	Af        Language = "af"
	AmET      Language = "am-ET"
	Am        Language = "am"
	ArAE      Language = "ar-AE"
	ArDZ      Language = "ar-DZ"
	ArYE      Language = "ar-YE"
	ArIQ      Language = "ar-IQ"
	ArEG      Language = "ar-EG"
	ArOM      Language = "ar-OM"
	ArQA      Language = "ar-QA"
	ArKW      Language = "ar-KW"
	ArSA      Language = "ar-SA"
	ArSY      Language = "ar-SY"
	ArTN      Language = "ar-TN"
	ArBH      Language = "ar-BH"
	ArMA      Language = "ar-MA"
	ArJO      Language = "ar-JO"
	ArLY      Language = "ar-LY"
	ArLB      Language = "ar-LB"
	Ar        Language = "ar"
	GswFR     Language = "gsw-FR"
	Gsw       Language = "gsw"
	SqAL      Language = "sq-AL"
	Sq        Language = "sq"
	HyAM      Language = "hy-AM"
	Hy        Language = "hy"
	ItIT      Language = "it-IT"
	ItCH      Language = "it-CH"
	It        Language = "it"
	IuCans    Language = "iu-Cans"
	IuCansCA  Language = "iu-Cans-CA"
	IuLatn    Language = "iu-Latn"
	IuLatnCA  Language = "iu-Latn-CA"
	Iu        Language = "iu"
	IgNG      Language = "ig-NG"
	Ig        Language = "ig"
	IdID      Language = "id-ID"
	Id        Language = "id"
	IiCN      Language = "ii-CN"
	Ii        Language = "ii"
	UgCN      Language = "ug-CN"
	Ug        Language = "ug"
	CyGB      Language = "cy-GB"
	Cy        Language = "cy"
	WoSN      Language = "wo-SN"
	Wo        Language = "wo"
	UkUA      Language = "uk-UA"
	Uk        Language = "uk"
	UzCyrl    Language = "uz-Cyrl"
	UzCyrlUZ  Language = "uz-Cyrl-UZ"
	UzLatn    Language = "uz-Latn"
	UzLatnUZ  Language = "uz-Latn-UZ"
	Uz        Language = "uz"
	UrPK      Language = "ur-PK"
	Ur        Language = "ur"
	EtEE      Language = "et-EE"
	Et        Language = "et"
	OcFR      Language = "oc-FR"
	Oc        Language = "oc"
	NlNL      Language = "nl-NL"
	NlBE      Language = "nl-BE"
	Nl        Language = "nl"
	OrIN      Language = "or-IN"
	Or        Language = "or"
	KkKZ      Language = "kk-KZ"
	Kk        Language = "kk"
	CaES      Language = "ca-ES"
	Ca        Language = "ca"
	KnIN      Language = "kn-IN"
	GlES      Language = "gl-ES"
	Gl        Language = "gl"
	Kn        Language = "kn"
	QutGT     Language = "qut-GT"
	Qut       Language = "qut"
	RwRW      Language = "rw-RW"
	Rw        Language = "rw"
	ElGR      Language = "el-GR"
	El        Language = "el"
	KyKG      Language = "ky-KG"
	Ky        Language = "ky"
	GuIN      Language = "gu-IN"
	Gu        Language = "gu"
	KmKH      Language = "km-KH"
	Km        Language = "km"
	KlGL      Language = "kl-GL"
	Kl        Language = "kl"
	KaGE      Language = "ka-GE"
	Ka        Language = "ka"
	HrHR      Language = "hr-HR"
	HrBA      Language = "hr-BA"
	Hr        Language = "hr"
	QuzEC     Language = "quz-EC"
	QuzPE     Language = "quz-PE"
	QuzBO     Language = "quz-BO"
	Quz       Language = "quz"
	Kok       Language = "kok"
	XhZA      Language = "xh-ZA"
	Xh        Language = "xh"
	CoFR      Language = "co-FR"
	Co        Language = "co"
	KokIN     Language = "kok-IN"
	Smn       Language = "smn"
	Sms       Language = "sms"
	Smj       Language = "smj"
	Sma       Language = "sma"
	Se        Language = "se"
	SmnFI     Language = "smn-FI"
	SmsFI     Language = "sms-FI"
	SmjSE     Language = "smj-SE"
	SmjNO     Language = "smj-NO"
	SmaSE     Language = "sma-SE"
	SmaNO     Language = "sma-NO"
	SeSE      Language = "se-SE"
	SeNO      Language = "se-NO"
	SeFI      Language = "se-FI"
	SaIN      Language = "sa-IN"
	Sa        Language = "sa"
	SyrSY     Language = "syr-SY"
	Syr       Language = "syr"
	SiLK      Language = "si-LK"
	Si        Language = "si"
	SvSE      Language = "sv-SE"
	SvFI      Language = "sv-FI"
	Sv        Language = "sv"
	ZuZA      Language = "zu-ZA"
	Zu        Language = "zu"
	GdGB      Language = "gd-GB"
	Gd        Language = "gd"
	EsAR      Language = "es-AR"
	EsUY      Language = "es-UY"
	EsEC      Language = "es-EC"
	EsSV      Language = "es-SV"
	EsGT      Language = "es-GT"
	EsCR      Language = "es-CR"
	EsCO      Language = "es-CO"
	EsES      Language = "es-ES"
	EsCL      Language = "es-CL"
	EsDO      Language = "es-DO"
	EsNI      Language = "es-NI"
	EsPA      Language = "es-PA"
	EsPY      Language = "es-PY"
	EsPR      Language = "es-PR"
	EsVE      Language = "es-VE"
	EsPE      Language = "es-PE"
	EsBO      Language = "es-BO"
	EsHN      Language = "es-HN"
	EsMX      Language = "es-MX"
	EsUS      Language = "es-US"
	Es        Language = "es"
	SkSK      Language = "sk-SK"
	Sk        Language = "sk"
	SlSI      Language = "sl-SI"
	Sl        Language = "sl"
	SwKE      Language = "sw-KE"
	Sw        Language = "sw"
	NsoZA     Language = "nso-ZA"
	Nso       Language = "nso"
	TnZA      Language = "tn-ZA"
	Tn        Language = "tn"
	SrCyrl    Language = "sr-Cyrl"
	SrCyrlRS  Language = "sr-Cyrl-RS"
	SrLatnCS  Language = "sr-Latn-CS"
	SrCyrlBA  Language = "sr-Cyrl-BA"
	SrCyrlME  Language = "sr-Cyrl-ME"
	SrLatn    Language = "sr-Latn"
	SrLatnRS  Language = "sr-Latn-RS"
	SrCyrlCS  Language = "sr-Cyrl-CS"
	SrLatnBA  Language = "sr-Latn-BA"
	SrLatnME  Language = "sr-Latn-ME"
	Sr        Language = "sr"
	ThTH      Language = "th-TH"
	Th        Language = "th"
	TgCyrl    Language = "tg-Cyrl"
	TgCyrlTJ  Language = "tg-Cyrl-TJ"
	Tg        Language = "tg"
	TtRU      Language = "tt-RU"
	Tt        Language = "tt"
	TzmLatn   Language = "tzm-Latn"
	TzmLatnDZ Language = "tzm-Latn-DZ"
	Tzm       Language = "tzm"
	TaIN      Language = "ta-IN"
	Ta        Language = "ta"
	Prs       Language = "prs"
	PrsAF     Language = "prs-AF"
	CsCZ      Language = "cs-CZ"
	Cs        Language = "cs"
	BoCN      Language = "bo-CN"
	Bo        Language = "bo"
	DvMV      Language = "dv-MV"
	Dv        Language = "dv"
	TeIN      Language = "te-IN"
	Te        Language = "te"
	DaDK      Language = "da-DK"
	Da        Language = "da"
	DeAT      Language = "de-AT"
	DeCH      Language = "de-CH"
	DeDE      Language = "de-DE"
	DeLI      Language = "de-LI"
	DeLU      Language = "de-LU"
	De        Language = "de"
	TkTM      Language = "tk-TM"
	Tk        Language = "tk"
	TrTR      Language = "tr-TR"
	Tr        Language = "tr"
	NeNP      Language = "ne-NP"
	Ne        Language = "ne"
	Nn        Language = "nn"
	Nb        Language = "nb"
	No        Language = "no"
	NnNO      Language = "nn-NO"
	NbNO      Language = "nb-NO"
	HaLatn    Language = "ha-Latn"
	HaLatnNG  Language = "ha-Latn-NG"
	Ha        Language = "ha"
	BaRU      Language = "ba-RU"
	Ba        Language = "ba"
	PsAF      Language = "ps-AF"
	Ps        Language = "ps"
	EuES      Language = "eu-ES"
	Eu        Language = "eu"
	HuHU      Language = "hu-HU"
	Hu        Language = "hu"
	PaIN      Language = "pa-IN"
	Pa        Language = "pa"
	HiIN      Language = "hi-IN"
	Hi        Language = "hi"
	FilPH     Language = "fil-PH"
	Fil       Language = "fil"
	FiFI      Language = "fi-FI"
	Fi        Language = "fi"
	FoFO      Language = "fo-FO"
	Fo        Language = "fo"
	FrCA      Language = "fr-CA"
	FrCH      Language = "fr-CH"
	FrFR      Language = "fr-FR"
	FrBE      Language = "fr-BE"
	FrMC      Language = "fr-MC"
	FrLU      Language = "fr-LU"
	Fr        Language = "fr"
	FyNL      Language = "fy-NL"
	Fy        Language = "fy"
	BgBG      Language = "bg-BG"
	Bg        Language = "bg"
	BrFR      Language = "br-FR"
	Br        Language = "br"
	ViVN      Language = "vi-VN"
	Vi        Language = "vi"
	HeIL      Language = "he-IL"
	He        Language = "he"
	BeBY      Language = "be-BY"
	Be        Language = "be"
	Fa        Language = "fa"
	FaIR      Language = "fa-IR"
	BnIN      Language = "bn-IN"
	BnBD      Language = "bn-BD"
	Bn        Language = "bn"
	PlPL      Language = "pl-PL"
	Pl        Language = "pl"
	BsCyrl    Language = "bs-Cyrl"
	BsCyrlBA  Language = "bs-Cyrl-BA"
	BsLatn    Language = "bs-Latn"
	BsLatnBA  Language = "bs-Latn-BA"
	Bs        Language = "bs"
	PtBR      Language = "pt-BR"
	PtPT      Language = "pt-PT"
	Pt        Language = "pt"
	MiNZ      Language = "mi-NZ"
	Mi        Language = "mi"
	Mk        Language = "mk"
	MkMK      Language = "mk-MK"
	ArnCL     Language = "arn-CL"
	Arn       Language = "arn"
	MrIN      Language = "mr-IN"
	Mr        Language = "mr"
	Ml        Language = "ml"
	MlIN      Language = "ml-IN"
	MtMT      Language = "mt-MT"
	Mt        Language = "mt"
	MsBN      Language = "ms-BN"
	MsMY      Language = "ms-MY"
	Ms        Language = "ms"
	MohCA     Language = "moh-CA"
	Moh       Language = "moh"
	MnCyrl    Language = "mn-Cyrl"
	MnMN      Language = "mn-MN"
	MnMong    Language = "mn-Mong"
	MnMongCN  Language = "mn-Mong-CN"
	Mn        Language = "mn"
	SahRU     Language = "sah-RU"
	Sah       Language = "sah"
	YoNG      Language = "yo-NG"
	Yo        Language = "yo"
	LoLA      Language = "lo-LA"
	Lo        Language = "lo"
	LvLV      Language = "lv-LV"
	Lv        Language = "lv"
	LtLT      Language = "lt-LT"
	Lt        Language = "lt"
	RoRO      Language = "ro-RO"
	Ro        Language = "ro"
	LbLU      Language = "lb-LU"
	Lb        Language = "lb"
	RuRU      Language = "ru-RU"
	Ru        Language = "ru"
	RmCH      Language = "rm-CH"
	Rm        Language = "rm"
	EnIE      Language = "en-IE"
	EnIN      Language = "en-IN"
	EnAU      Language = "en-AU"
	EnCA      Language = "en-CA"
	En029     Language = "en-029"
	EnJM      Language = "en-JM"
	EnSG      Language = "en-SG"
	EnZW      Language = "en-ZW"
	EnTT      Language = "en-TT"
	EnNZ      Language = "en-NZ"
	EnPH      Language = "en-PH"
	EnBZ      Language = "en-BZ"
	EnMY      Language = "en-MY"
	EnGB      Language = "en-GB"
	EnZA      Language = "en-ZA"
	EnUS      Language = "en-US"
	En        Language = "en"
	DsbDE     Language = "dsb-DE"
	Dsb       Language = "dsb"
	KoKR      Language = "ko-KR"
	Ko        Language = "ko"
	HsbDE     Language = "hsb-DE"
	Hsb       Language = "hsb"
	ZhHans    Language = "zh-Hans"
	ZhSG      Language = "zh-SG"
	ZhCN      Language = "zh-CN"
	ZhHant    Language = "zh-Hant"
	ZhHK      Language = "zh-HK"
	ZhTW      Language = "zh-TW"
	ZhMO      Language = "zh-MO"
	Zh        Language = "zh"
	JaJP      Language = "ja-JP"
	Ja        Language = "ja"
)

func (p Language) String() string {
	return string(p)
}

func (p Language) Ptr() *Language {
	return &p
}

type LanguageData struct {
	Code     Language
	Template string
	Decimal  string
	Thousand string
}
type Languages map[Language]*LanguageData

var languages = Languages{
	IsIS:      {Code: IsIS, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Is:        {Code: Is, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	GaIE:      {Code: GaIE, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Ga:        {Code: Ga, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	AzCyrl:    {Code: AzCyrl, Decimal: ",", Thousand: ".", Template: "%d %s"},   // 1.234,12 US$
	AzCyrlAZ:  {Code: AzCyrlAZ, Decimal: ",", Thousand: ".", Template: "%d %s"}, // 1.234,12 US$
	AzLatn:    {Code: AzLatn, Decimal: ".", Thousand: ",", Template: "%s %d"},   // $ 1,234.12
	AzLatnAZ:  {Code: AzLatnAZ, Decimal: ".", Thousand: ",", Template: "%s %d"}, // $ 1,234.12
	Az:        {Code: Az, Decimal: ".", Thousand: ",", Template: "%s %d"},       // $ 1,234.12
	AsIN:      {Code: AsIN, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	As:        {Code: As, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	AfZA:      {Code: AfZA, Decimal: ",", Thousand: " ", Template: "%s %d"},     // USD 1 234,12
	Af:        {Code: Af, Decimal: ",", Thousand: " ", Template: "%s %d"},       // USD 1 234,12
	AmET:      {Code: AmET, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	Am:        {Code: Am, Decimal: ".", Thousand: ",", Template: "%d%s"},        // US$1,234.12
	ArAE:      {Code: ArAE, Decimal: ".", Thousand: ",", Template: "%d %s"},     // ‏1,234.12 US$
	ArDZ:      {Code: ArDZ, Decimal: ",", Thousand: ".", Template: "%d %s"},     // ‏1.234,12 US$
	ArYE:      {Code: ArYE, Decimal: ",", Thousand: ".", Template: "%d %s"},     // ‏١٬٢٣٤٫١٢ US$
	ArIQ:      {Code: ArIQ, Decimal: ",", Thousand: ".", Template: "%d %s"},     // ‏١٬٢٣٤٫١٢ US$
	ArEG:      {Code: ArEG, Decimal: ",", Thousand: ".", Template: "%d %s"},     // ‏١٬٢٣٤٫١٢ US$
	ArOM:      {Code: ArOM, Decimal: ",", Thousand: ".", Template: "%d %s"},     // ‏١٬٢٣٤٫١٢ US$
	ArQA:      {Code: ArQA, Decimal: ",", Thousand: ".", Template: "%d %s"},     // ‏١٬٢٣٤٫١٢ US$
	ArKW:      {Code: ArKW, Decimal: ",", Thousand: ".", Template: "%d %s"},     // ‏١٬٢٣٤٫١٢ US$
	ArSA:      {Code: ArSA, Decimal: ",", Thousand: ".", Template: "%d %s"},     // ‏١٬٢٣٤٫١٢ US$
	ArSY:      {Code: ArSY, Decimal: ",", Thousand: ".", Template: "%d %s"},     // ‏١٬٢٣٤٫١٢ US$
	ArTN:      {Code: ArTN, Decimal: ",", Thousand: ".", Template: "%d %s"},     // ‏1.234,12 US$
	ArBH:      {Code: ArBH, Decimal: ",", Thousand: ".", Template: "%d %s"},     // ‏١٬٢٣٤٫١٢ US$
	ArMA:      {Code: ArMA, Decimal: ",", Thousand: ".", Template: "%d %s"},     // ‏1.234,12 US$
	ArJO:      {Code: ArJO, Decimal: ",", Thousand: ".", Template: "%d %s"},     // ‏١٬٢٣٤٫١٢ US$
	ArLY:      {Code: ArLY, Decimal: ",", Thousand: ".", Template: "%d %s"},     // ‏1.234,12 US$
	ArLB:      {Code: ArLB, Decimal: ",", Thousand: ".", Template: "%d %s"},     // ‏١٬٢٣٤٫١٢ US$
	Ar:        {Code: Ar, Decimal: ".", Thousand: ",", Template: "%d %s"},       // ‏1,234.12 US$
	GswFR:     {Code: GswFR, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	Gsw:       {Code: Gsw, Decimal: ".", Thousand: ",", Template: "%s%d"},       // $1,234.12
	SqAL:      {Code: SqAL, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Sq:        {Code: Sq, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	HyAM:      {Code: HyAM, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Hy:        {Code: Hy, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	ItIT:      {Code: ItIT, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1.234,12 USD
	ItCH:      {Code: ItCH, Decimal: ".", Thousand: "’", Template: "%s %d"},     // USD 1’234.12
	It:        {Code: It, Decimal: ",", Thousand: ".", Template: "%d %s"},       // 1.234,12 USD
	IuCans:    {Code: IuCans, Decimal: ".", Thousand: ",", Template: "%s%d"},    // $1,234.12
	IuCansCA:  {Code: IuCansCA, Decimal: ".", Thousand: ",", Template: "%s%d"},  // $1,234.12
	IuLatn:    {Code: IuLatn, Decimal: ".", Thousand: ",", Template: "%s%d"},    // $1,234.12
	IuLatnCA:  {Code: IuLatnCA, Decimal: ".", Thousand: ",", Template: "%s%d"},  // $1,234.12
	Iu:        {Code: Iu, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	IgNG:      {Code: IgNG, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Ig:        {Code: Ig, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	IdID:      {Code: IdID, Decimal: ",", Thousand: ".", Template: "%d%s"},      // US$1.234,12
	Id:        {Code: Id, Decimal: ",", Thousand: ".", Template: "%d%s"},        // US$1.234,12
	IiCN:      {Code: IiCN, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Ii:        {Code: Ii, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	UgCN:      {Code: UgCN, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Ug:        {Code: Ug, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	CyGB:      {Code: CyGB, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Cy:        {Code: Cy, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	WoSN:      {Code: WoSN, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Wo:        {Code: Wo, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	UkUA:      {Code: UkUA, Decimal: ",", Thousand: " ", Template: "%d %s"},     // 1 234,12 USD
	Uk:        {Code: Uk, Decimal: ",", Thousand: " ", Template: "%d %s"},       // 1 234,12 USD
	UzCyrl:    {Code: UzCyrl, Decimal: ",", Thousand: " ", Template: "%d %s"},   // 1 234,12 US$
	UzCyrlUZ:  {Code: UzCyrlUZ, Decimal: ",", Thousand: " ", Template: "%d %s"}, // 1 234,12 US$
	UzLatn:    {Code: UzLatn, Decimal: ".", Thousand: ",", Template: "%s %d"},   // $ 1,234.12
	UzLatnUZ:  {Code: UzLatnUZ, Decimal: ".", Thousand: ",", Template: "%s %d"}, // $ 1,234.12
	Uz:        {Code: Uz, Decimal: ".", Thousand: ",", Template: "%s %d"},       // $ 1,234.12
	UrPK:      {Code: UrPK, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Ur:        {Code: Ur, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	EtEE:      {Code: EtEE, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1234,12 $
	Et:        {Code: Et, Decimal: ",", Thousand: ".", Template: "%d %s"},       // 1234,12 $
	OcFR:      {Code: OcFR, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Oc:        {Code: Oc, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	NlNL:      {Code: NlNL, Decimal: ",", Thousand: ".", Template: "%d%s"},      // US$ 1.234,12
	NlBE:      {Code: NlBE, Decimal: ",", Thousand: ".", Template: "%d%s"},      // US$ 1.234,12
	Nl:        {Code: Nl, Decimal: ",", Thousand: ".", Template: "%d%s"},        // US$ 1.234,12
	OrIN:      {Code: OrIN, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Or:        {Code: Or, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	KkKZ:      {Code: KkKZ, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Kk:        {Code: Kk, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	CaES:      {Code: CaES, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1.234,12 USD
	Ca:        {Code: Ca, Decimal: ",", Thousand: ".", Template: "%d %s"},       // 1.234,12 USD
	KnIN:      {Code: KnIN, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	GlES:      {Code: GlES, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Gl:        {Code: Gl, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	Kn:        {Code: Kn, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	QutGT:     {Code: QutGT, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	Qut:       {Code: Qut, Decimal: ".", Thousand: ",", Template: "%s%d"},       // $1,234.12
	RwRW:      {Code: RwRW, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Rw:        {Code: Rw, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	ElGR:      {Code: ElGR, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1.234,12 $
	El:        {Code: El, Decimal: ",", Thousand: ".", Template: "%d %s"},       // 1.234,12 $
	KyKG:      {Code: KyKG, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Ky:        {Code: Ky, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	GuIN:      {Code: GuIN, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	Gu:        {Code: Gu, Decimal: ".", Thousand: ",", Template: "%d%s"},        // US$1,234.12
	KmKH:      {Code: KmKH, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Km:        {Code: Km, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	KlGL:      {Code: KlGL, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Kl:        {Code: Kl, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	KaGE:      {Code: KaGE, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Ka:        {Code: Ka, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	HrHR:      {Code: HrHR, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1.234,12 USD
	HrBA:      {Code: HrBA, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1.234,12 USD
	Hr:        {Code: Hr, Decimal: ",", Thousand: ".", Template: "%d %s"},       // 1.234,12 USD
	QuzEC:     {Code: QuzEC, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	QuzPE:     {Code: QuzPE, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	QuzBO:     {Code: QuzBO, Decimal: ",", Thousand: ".", Template: "%s %d"},    // $ 1.234,12
	Quz:       {Code: Quz, Decimal: ".", Thousand: ",", Template: "%s%d"},       // $1,234.12
	Kok:       {Code: Kok, Decimal: ".", Thousand: ",", Template: "%s%d"},       // $1,234.12
	XhZA:      {Code: XhZA, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Xh:        {Code: Xh, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	CoFR:      {Code: CoFR, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Co:        {Code: Co, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	KokIN:     {Code: KokIN, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	Smn:       {Code: Smn, Decimal: ".", Thousand: ",", Template: "%s%d"},       // $1,234.12
	Sms:       {Code: Sms, Decimal: ".", Thousand: ",", Template: "%s%d"},       // $1,234.12
	Smj:       {Code: Smj, Decimal: ".", Thousand: ",", Template: "%s%d"},       // $1,234.12
	Sma:       {Code: Sma, Decimal: ".", Thousand: ",", Template: "%s%d"},       // $1,234.12
	Se:        {Code: Se, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	SmnFI:     {Code: SmnFI, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	SmsFI:     {Code: SmsFI, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	SmjSE:     {Code: SmjSE, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	SmjNO:     {Code: SmjNO, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	SmaSE:     {Code: SmaSE, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	SmaNO:     {Code: SmaNO, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	SeSE:      {Code: SeSE, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	SeNO:      {Code: SeNO, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	SeFI:      {Code: SeFI, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	SaIN:      {Code: SaIN, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Sa:        {Code: Sa, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	SyrSY:     {Code: SyrSY, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	Syr:       {Code: Syr, Decimal: ".", Thousand: ",", Template: "%s%d"},       // $1,234.12
	SiLK:      {Code: SiLK, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Si:        {Code: Si, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	SvSE:      {Code: SvSE, Decimal: ",", Thousand: " ", Template: "%d %s"},     // 1 234,12 US$
	SvFI:      {Code: SvFI, Decimal: ",", Thousand: " ", Template: "%d %s"},     // 1 234,12 US$
	Sv:        {Code: Sv, Decimal: ",", Thousand: " ", Template: "%d %s"},       // 1 234,12 US$
	ZuZA:      {Code: ZuZA, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Zu:        {Code: Zu, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	GdGB:      {Code: GdGB, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Gd:        {Code: Gd, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	EsAR:      {Code: EsAR, Decimal: ",", Thousand: ".", Template: "%d%s"},      // US$ 1.234,12
	EsUY:      {Code: EsUY, Decimal: ",", Thousand: ".", Template: "%d%s"},      // US$ 1.234,12
	EsEC:      {Code: EsEC, Decimal: ",", Thousand: ".", Template: "%s%d"},      // $1.234,12
	EsSV:      {Code: EsSV, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	EsGT:      {Code: EsGT, Decimal: ".", Thousand: ",", Template: "%s %d"},     // USD 1,234.12
	EsCR:      {Code: EsCR, Decimal: ",", Thousand: " ", Template: "%s %d"},     // USD 1 234,12
	EsCO:      {Code: EsCO, Decimal: ",", Thousand: ".", Template: "%d%s"},      // US$ 1.234,12
	EsES:      {Code: EsES, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1234,12 US$
	EsCL:      {Code: EsCL, Decimal: ",", Thousand: ".", Template: "%d%s"},      // US$1.234,12
	EsDO:      {Code: EsDO, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	EsNI:      {Code: EsNI, Decimal: ".", Thousand: ",", Template: "%s %d"},     // USD 1,234.12
	EsPA:      {Code: EsPA, Decimal: ".", Thousand: ",", Template: "%s %d"},     // USD 1,234.12
	EsPY:      {Code: EsPY, Decimal: ",", Thousand: ".", Template: "%s %d"},     // USD 1.234,12
	EsPR:      {Code: EsPR, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	EsVE:      {Code: EsVE, Decimal: ",", Thousand: ".", Template: "%s %d"},     // USD 1.234,12
	EsPE:      {Code: EsPE, Decimal: ".", Thousand: ",", Template: "%s %d"},     // USD 1,234.12
	EsBO:      {Code: EsBO, Decimal: ",", Thousand: ".", Template: "%s %d"},     // USD 1.234,12
	EsHN:      {Code: EsHN, Decimal: ".", Thousand: ",", Template: "%s %d"},     // USD 1,234.12
	EsMX:      {Code: EsMX, Decimal: ".", Thousand: ",", Template: "%s %d"},     // USD 1,234.12
	EsUS:      {Code: EsUS, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Es:        {Code: Es, Decimal: ",", Thousand: ".", Template: "%d %s"},       // 1234,12 US$
	SkSK:      {Code: SkSK, Decimal: ",", Thousand: " ", Template: "%d %s"},     // 1 234,12 USD
	Sk:        {Code: Sk, Decimal: ",", Thousand: " ", Template: "%d %s"},       // 1 234,12 USD
	SlSI:      {Code: SlSI, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1.234,12 $
	Sl:        {Code: Sl, Decimal: ",", Thousand: ".", Template: "%d %s"},       // 1.234,12 $
	SwKE:      {Code: SwKE, Decimal: ".", Thousand: ",", Template: "%s %d"},     // $ 1,234.12
	Sw:        {Code: Sw, Decimal: ".", Thousand: ",", Template: "%d%s"},        // US$ 1,234.12
	NsoZA:     {Code: NsoZA, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	Nso:       {Code: Nso, Decimal: ".", Thousand: ",", Template: "%s%d"},       // $1,234.12
	TnZA:      {Code: TnZA, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Tn:        {Code: Tn, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	SrCyrl:    {Code: SrCyrl, Decimal: ",", Thousand: ".", Template: "%d %s"},   // 1.234,12 US$
	SrCyrlRS:  {Code: SrCyrlRS, Decimal: ",", Thousand: ".", Template: "%d %s"}, // 1.234,12 US$
	SrLatnCS:  {Code: SrLatnCS, Decimal: ",", Thousand: ".", Template: "%d %s"}, // 1.234,12 US$
	SrCyrlBA:  {Code: SrCyrlBA, Decimal: ",", Thousand: ".", Template: "%d %s"}, // 1.234,12 US$
	SrCyrlME:  {Code: SrCyrlME, Decimal: ",", Thousand: ".", Template: "%d %s"}, // 1.234,12 US$
	SrLatn:    {Code: SrLatn, Decimal: ",", Thousand: ".", Template: "%d %s"},   // 1.234,12 US$
	SrLatnRS:  {Code: SrLatnRS, Decimal: ",", Thousand: ".", Template: "%d %s"}, // 1.234,12 US$
	SrCyrlCS:  {Code: SrCyrlCS, Decimal: ",", Thousand: ".", Template: "%d %s"}, // 1.234,12 US$
	SrLatnBA:  {Code: SrLatnBA, Decimal: ",", Thousand: ".", Template: "%d %s"}, // 1.234,12 US$
	SrLatnME:  {Code: SrLatnME, Decimal: ",", Thousand: ".", Template: "%d %s"}, // 1.234,12 US$
	Sr:        {Code: Sr, Decimal: ",", Thousand: ".", Template: "%d %s"},       // 1.234,12 US$
	ThTH:      {Code: ThTH, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	Th:        {Code: Th, Decimal: ".", Thousand: ",", Template: "%d%s"},        // US$1,234.12
	TgCyrl:    {Code: TgCyrl, Decimal: ".", Thousand: ",", Template: "%s%d"},    // $1,234.12
	TgCyrlTJ:  {Code: TgCyrlTJ, Decimal: ".", Thousand: ",", Template: "%s%d"},  // $1,234.12
	Tg:        {Code: Tg, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	TtRU:      {Code: TtRU, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Tt:        {Code: Tt, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	TzmLatn:   {Code: TzmLatn, Decimal: ".", Thousand: ",", Template: "%s%d"},   // $1,234.12
	TzmLatnDZ: {Code: TzmLatnDZ, Decimal: ".", Thousand: ",", Template: "%s%d"}, // $1,234.12
	Tzm:       {Code: Tzm, Decimal: ".", Thousand: ",", Template: "%s%d"},       // $1,234.12
	TaIN:      {Code: TaIN, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Ta:        {Code: Ta, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	Prs:       {Code: Prs, Decimal: ",", Thousand: ".", Template: "%s %d"},      // $ ۱٬۲۳۴٫۱۲
	PrsAF:     {Code: PrsAF, Decimal: ",", Thousand: ".", Template: "%s %d"},    // $ ۱٬۲۳۴٫۱۲
	CsCZ:      {Code: CsCZ, Decimal: ",", Thousand: " ", Template: "%d %s"},     // 1 234,12 US$
	Cs:        {Code: Cs, Decimal: ",", Thousand: " ", Template: "%d %s"},       // 1 234,12 US$
	BoCN:      {Code: BoCN, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Bo:        {Code: Bo, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	DvMV:      {Code: DvMV, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Dv:        {Code: Dv, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	TeIN:      {Code: TeIN, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Te:        {Code: Te, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	DaDK:      {Code: DaDK, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1.234,12 US$
	Da:        {Code: Da, Decimal: ",", Thousand: ".", Template: "%d %s"},       // 1.234,12 US$
	DeAT:      {Code: DeAT, Decimal: ",", Thousand: ".", Template: "%s %d"},     // $ 1.234,12
	DeCH:      {Code: DeCH, Decimal: ".", Thousand: "’", Template: "%s %d"},     // $ 1’234.12
	DeDE:      {Code: DeDE, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1.234,12 $
	DeLI:      {Code: DeLI, Decimal: ".", Thousand: "’", Template: "%s %d"},     // $ 1’234.12
	DeLU:      {Code: DeLU, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1.234,12 $
	De:        {Code: De, Decimal: ",", Thousand: ".", Template: "%d %s"},       // 1.234,12 $
	TkTM:      {Code: TkTM, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Tk:        {Code: Tk, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	TrTR:      {Code: TrTR, Decimal: ",", Thousand: ".", Template: "%s%d"},      // $1.234,12
	Tr:        {Code: Tr, Decimal: ",", Thousand: ".", Template: "%s%d"},        // $1.234,12
	NeNP:      {Code: NeNP, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Ne:        {Code: Ne, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	Nn:        {Code: Nn, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	Nb:        {Code: Nb, Decimal: ",", Thousand: " ", Template: "%s %d"},       // USD 1 234,12
	No:        {Code: No, Decimal: ",", Thousand: " ", Template: "%s %d"},       // USD 1 234,12
	NnNO:      {Code: NnNO, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	NbNO:      {Code: NbNO, Decimal: ",", Thousand: " ", Template: "%s %d"},     // USD 1 234,12
	HaLatn:    {Code: HaLatn, Decimal: ".", Thousand: ",", Template: "%s%d"},    // $1,234.12
	HaLatnNG:  {Code: HaLatnNG, Decimal: ".", Thousand: ",", Template: "%s%d"},  // $1,234.12
	Ha:        {Code: Ha, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	BaRU:      {Code: BaRU, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Ba:        {Code: Ba, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	PsAF:      {Code: PsAF, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Ps:        {Code: Ps, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	EuES:      {Code: EuES, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Eu:        {Code: Eu, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	HuHU:      {Code: HuHU, Decimal: ",", Thousand: " ", Template: "%d %s"},     // 1 234,12 USD
	Hu:        {Code: Hu, Decimal: ",", Thousand: " ", Template: "%d %s"},       // 1 234,12 USD
	PaIN:      {Code: PaIN, Decimal: ".", Thousand: ",", Template: "%s %d"},     // $ 1,234.12
	Pa:        {Code: Pa, Decimal: ".", Thousand: ",", Template: "%s %d"},       // $ 1,234.12
	HiIN:      {Code: HiIN, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Hi:        {Code: Hi, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	FilPH:     {Code: FilPH, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	Fil:       {Code: Fil, Decimal: ".", Thousand: ",", Template: "%s%d"},       // $1,234.12
	FiFI:      {Code: FiFI, Decimal: ",", Thousand: " ", Template: "%d %s"},     // 1 234,12 $
	Fi:        {Code: Fi, Decimal: ",", Thousand: " ", Template: "%d %s"},       // 1 234,12 $
	FoFO:      {Code: FoFO, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Fo:        {Code: Fo, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	FrCA:      {Code: FrCA, Decimal: ",", Thousand: " ", Template: "%d %s"},     // 1 234,12 $ US
	FrCH:      {Code: FrCH, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1 234.12 $US
	FrFR:      {Code: FrFR, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1 234,12 $US
	FrBE:      {Code: FrBE, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1 234,12 $US
	FrMC:      {Code: FrMC, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1 234,12 $US
	FrLU:      {Code: FrLU, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1.234,12 $US
	Fr:        {Code: Fr, Decimal: ",", Thousand: ".", Template: "%d %s"},       // 1 234,12 $US
	FyNL:      {Code: FyNL, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Fy:        {Code: Fy, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	BgBG:      {Code: BgBG, Decimal: ",", Thousand: ".", Template: "%s%d"},      // 1234,12 щ.д.
	Bg:        {Code: Bg, Decimal: ",", Thousand: ".", Template: "%s%d"},        // 1234,12 щ.д.
	BrFR:      {Code: BrFR, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Br:        {Code: Br, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	ViVN:      {Code: ViVN, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1.234,12 US$
	Vi:        {Code: Vi, Decimal: ",", Thousand: ".", Template: "%d %s"},       // 1.234,12 US$
	HeIL:      {Code: HeIL, Decimal: ".", Thousand: ",", Template: "%d%s"},      // ‏1,234.12 ‏$
	He:        {Code: He, Decimal: ".", Thousand: ",", Template: "%d%s"},        // ‏1,234.12 ‏$
	BeBY:      {Code: BeBY, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Be:        {Code: Be, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	Fa:        {Code: Fa, Decimal: ",", Thousand: ".", Template: "%d%s"},        // ‎$۱٬۲۳۴٫۱۲
	FaIR:      {Code: FaIR, Decimal: ",", Thousand: ".", Template: "%d%s"},      // ‎$۱٬۲۳۴٫۱۲
	BnIN:      {Code: BnIN, Decimal: ",", Thousand: ".", Template: "%s%d"},      // $১,২৩৪.১২
	BnBD:      {Code: BnBD, Decimal: ",", Thousand: ".", Template: "%d %s"},     // ১,২৩৪.১২ US$
	Bn:        {Code: Bn, Decimal: ",", Thousand: ".", Template: "%d %s"},       // ১,২৩৪.১২ US$
	PlPL:      {Code: PlPL, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1234,12 USD
	Pl:        {Code: Pl, Decimal: ",", Thousand: ".", Template: "%d %s"},       // 1234,12 USD
	BsCyrl:    {Code: BsCyrl, Decimal: ",", Thousand: ".", Template: "%d %s"},   // 1.234,12 US$
	BsCyrlBA:  {Code: BsCyrlBA, Decimal: ",", Thousand: ".", Template: "%d %s"}, // 1.234,12 US$
	BsLatn:    {Code: BsLatn, Decimal: ".", Thousand: ",", Template: "%s %d"},   // $ 1,234.12
	BsLatnBA:  {Code: BsLatnBA, Decimal: ".", Thousand: ",", Template: "%s %d"}, // $ 1,234.12
	Bs:        {Code: Bs, Decimal: ".", Thousand: ",", Template: "%s %d"},       // $ 1,234.12
	PtBR:      {Code: PtBR, Decimal: ",", Thousand: ".", Template: "%d%s"},      // US$ 1.234,12
	PtPT:      {Code: PtPT, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1234,12 US$
	Pt:        {Code: Pt, Decimal: ",", Thousand: ".", Template: "%d%s"},        // US$ 1.234,12
	MiNZ:      {Code: MiNZ, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Mi:        {Code: Mi, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	Mk:        {Code: Mk, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	MkMK:      {Code: MkMK, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	ArnCL:     {Code: ArnCL, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	Arn:       {Code: Arn, Decimal: ".", Thousand: ",", Template: "%s%d"},       // $1,234.12
	MrIN:      {Code: MrIN, Decimal: ",", Thousand: ".", Template: "%s%d"},      // $१,२३४.१२
	Mr:        {Code: Mr, Decimal: ",", Thousand: ".", Template: "%s%d"},        // $१,२३४.१२
	Ml:        {Code: Ml, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	MlIN:      {Code: MlIN, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	MtMT:      {Code: MtMT, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Mt:        {Code: Mt, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	MsBN:      {Code: MsBN, Decimal: ",", Thousand: ".", Template: "%s %d"},     // USD 1.234,12
	MsMY:      {Code: MsMY, Decimal: ".", Thousand: ",", Template: "%s %d"},     // USD 1,234.12
	Ms:        {Code: Ms, Decimal: ".", Thousand: ",", Template: "%s %d"},       // USD 1,234.12
	MohCA:     {Code: MohCA, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	Moh:       {Code: Moh, Decimal: ".", Thousand: ",", Template: "%s%d"},       // $1,234.12
	MnCyrl:    {Code: MnCyrl, Decimal: ".", Thousand: ",", Template: "%s%d"},    // $1,234.12
	MnMN:      {Code: MnMN, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	MnMong:    {Code: MnMong, Decimal: ".", Thousand: ",", Template: "%s%d"},    // $1,234.12
	MnMongCN:  {Code: MnMongCN, Decimal: ".", Thousand: ",", Template: "%s%d"},  // $1,234.12
	Mn:        {Code: Mn, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	SahRU:     {Code: SahRU, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	Sah:       {Code: Sah, Decimal: ".", Thousand: ",", Template: "%s%d"},       // $1,234.12
	YoNG:      {Code: YoNG, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Yo:        {Code: Yo, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	LoLA:      {Code: LoLA, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Lo:        {Code: Lo, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	LvLV:      {Code: LvLV, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1234,12 $
	Lv:        {Code: Lv, Decimal: ",", Thousand: ".", Template: "%d %s"},       // 1234,12 $
	LtLT:      {Code: LtLT, Decimal: ",", Thousand: " ", Template: "%d %s"},     // 1 234,12 USD
	Lt:        {Code: Lt, Decimal: ",", Thousand: " ", Template: "%d %s"},       // 1 234,12 USD
	RoRO:      {Code: RoRO, Decimal: ",", Thousand: ".", Template: "%d %s"},     // 1.234,12 USD
	Ro:        {Code: Ro, Decimal: ",", Thousand: ".", Template: "%d %s"},       // 1.234,12 USD
	LbLU:      {Code: LbLU, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Lb:        {Code: Lb, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	RuRU:      {Code: RuRU, Decimal: ",", Thousand: " ", Template: "%d %s"},     // 1 234,12 $
	Ru:        {Code: Ru, Decimal: ",", Thousand: " ", Template: "%d %s"},       // 1 234,12 $
	RmCH:      {Code: RmCH, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Rm:        {Code: Rm, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	EnIE:      {Code: EnIE, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	EnIN:      {Code: EnIN, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	EnAU:      {Code: EnAU, Decimal: ".", Thousand: ",", Template: "%s %d"},     // USD 1,234.12
	EnCA:      {Code: EnCA, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	En029:     {Code: En029, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	EnJM:      {Code: EnJM, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	EnSG:      {Code: EnSG, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	EnZW:      {Code: EnZW, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	EnTT:      {Code: EnTT, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	EnNZ:      {Code: EnNZ, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	EnPH:      {Code: EnPH, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	EnBZ:      {Code: EnBZ, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	EnMY:      {Code: EnMY, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	EnGB:      {Code: EnGB, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	EnZA:      {Code: EnZA, Decimal: ",", Thousand: " ", Template: "%d%s"},      // US$1 234,12
	EnUS:      {Code: EnUS, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	En:        {Code: En, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
	DsbDE:     {Code: DsbDE, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	Dsb:       {Code: Dsb, Decimal: ".", Thousand: ",", Template: "%s%d"},       // $1,234.12
	KoKR:      {Code: KoKR, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	Ko:        {Code: Ko, Decimal: ".", Thousand: ",", Template: "%d%s"},        // US$1,234.12
	HsbDE:     {Code: HsbDE, Decimal: ".", Thousand: ",", Template: "%s%d"},     // $1,234.12
	Hsb:       {Code: Hsb, Decimal: ".", Thousand: ",", Template: "%s%d"},       // $1,234.12
	ZhHans:    {Code: ZhHans, Decimal: ".", Thousand: ",", Template: "%d%s"},    // US$1,234.12
	ZhSG:      {Code: ZhSG, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	ZhCN:      {Code: ZhCN, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	ZhHant:    {Code: ZhHant, Decimal: ".", Thousand: ",", Template: "%d%s"},    // US$1,234.12
	ZhHK:      {Code: ZhHK, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	ZhTW:      {Code: ZhTW, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	ZhMO:      {Code: ZhMO, Decimal: ".", Thousand: ",", Template: "%d%s"},      // US$1,234.12
	Zh:        {Code: Zh, Decimal: ".", Thousand: ",", Template: "%d%s"},        // US$1,234.12
	JaJP:      {Code: JaJP, Decimal: ".", Thousand: ",", Template: "%s%d"},      // $1,234.12
	Ja:        {Code: Ja, Decimal: ".", Thousand: ",", Template: "%s%d"},        // $1,234.12
}

func (c Languages) LanguageByCode(code Language) *LanguageData {
	sc, ok := c[code]
	if !ok {
		return nil
	}

	return sc
}

func GetLanguage(code Language) *LanguageData {
	return languages.LanguageByCode(Language(code))
}
