package airtable

import "encoding/json"

type TimeZone int64

const (
	TimeZoneUTC TimeZone = iota
	TimeZoneClient
	TimeZoneAfricaAbidjan
	TimeZoneAfricaAccra
	TimeZoneAfricaAddisAbaba
	TimeZoneAfricaAlgiers
	TimeZoneAfricaAsmara
	TimeZoneAfricaBamako
	TimeZoneAfricaBangui
	TimeZoneAfricaBanjul
	TimeZoneAfricaBissau
	TimeZoneAfricaBlantyre
	TimeZoneAfricaBrazzaville
	TimeZoneAfricaBujumbura
	TimeZoneAfricaCairo
	TimeZoneAfricaCasablanca
	TimeZoneAfricaCeuta
	TimeZoneAfricaConakry
	TimeZoneAfricaDakar
	TimeZoneAfricaDaresSalaam
	TimeZoneAfricaDjibouti
	TimeZoneAfricaDouala
	TimeZoneAfricaElAaiun
	TimeZoneAfricaFreetown
	TimeZoneAfricaGaborone
	TimeZoneAfricaHarare
	TimeZoneAfricaJohannesburg
	TimeZoneAfricaJuba
	TimeZoneAfricaKampala
	TimeZoneAfricaKhartoum
	TimeZoneAfricaKigali
	TimeZoneAfricaKinshasa
	TimeZoneAfricaLagos
	TimeZoneAfricaLibreville
	TimeZoneAfricaLome
	TimeZoneAfricaLuanda
	TimeZoneAfricaLubumbashi
	TimeZoneAfricaLusaka
	TimeZoneAfricaMalabo
	TimeZoneAfricaMaputo
	TimeZoneAfricaMaseru
	TimeZoneAfricaMbabane
	TimeZoneAfricaMogadishu
	TimeZoneAfricaMonrovia
	TimeZoneAfricaNairobi
	TimeZoneAfricaNdjamena
	TimeZoneAfricaNiamey
	TimeZoneAfricaNouakchott
	TimeZoneAfricaOuagadougou
	TimeZoneAfricaPortoNovo
	TimeZoneAfricaSaoTome
	TimeZoneAfricaTripoli
	TimeZoneAfricaTunis
	TimeZoneAfricaWindhoek
	TimeZoneAmericaAdak
	TimeZoneAmericaAnchorage
	TimeZoneAmericaAnguilla
	TimeZoneAmericaAntigua
	TimeZoneAmericaAraguaina
	TimeZoneAmericaArgentinaBuenosAires
	TimeZoneAmericaArgentinaCatamarca
	TimeZoneAmericaArgentinaCordoba
	TimeZoneAmericaArgentinaJujuy
	TimeZoneAmericaArgentinaLaRioja
	TimeZoneAmericaArgentinaMendoza
	TimeZoneAmericaArgentinaRioGallegos
	TimeZoneAmericaArgentinaSalta
	TimeZoneAmericaArgentinaSanJuan
	TimeZoneAmericaArgentinaSanLuis
	TimeZoneAmericaArgentinaTucuman
	TimeZoneAmericaArgentinaUshuaia
	TimeZoneAmericaAruba
	TimeZoneAmericaAsuncion
	TimeZoneAmericaAtikokan
	TimeZoneAmericaBahia
	TimeZoneAmericaBahiaBanderas
	TimeZoneAmericaBarbados
	TimeZoneAmericaBelem
	TimeZoneAmericaBelize
	TimeZoneAmericaBlancSablon
	TimeZoneAmericaBoaVista
	TimeZoneAmericaBogota
	TimeZoneAmericaBoise
	TimeZoneAmericaCambridgeBay
	TimeZoneAmericaCampoGrande
	TimeZoneAmericaCancun
	TimeZoneAmericaCaracas
	TimeZoneAmericaCayenne
	TimeZoneAmericaCayman
	TimeZoneAmericaChicago
	TimeZoneAmericaChihuahua
	TimeZoneAmericaCostaRica
	TimeZoneAmericaCreston
	TimeZoneAmericaCuiaba
	TimeZoneAmericaCuracao
	TimeZoneAmericaDanmarkshavn
	TimeZoneAmericaDawson
	TimeZoneAmericaDawsonCreek
	TimeZoneAmericaDenver
	TimeZoneAmericaDetroit
	TimeZoneAmericaDominica
	TimeZoneAmericaEdmonton
	TimeZoneAmericaEirunepe
	TimeZoneAmericaElSalvador
	TimeZoneAmericaFortNelson
	TimeZoneAmericaFortaleza
	TimeZoneAmericaGlaceBay
	TimeZoneAmericaGodthab
	TimeZoneAmericaGooseBay
	TimeZoneAmericaGrandTurk
	TimeZoneAmericaGrenada
	TimeZoneAmericaGuadeloupe
	TimeZoneAmericaGuatemala
	TimeZoneAmericaGuayaquil
	TimeZoneAmericaGuyana
	TimeZoneAmericaHalifax
	TimeZoneAmericaHavana
	TimeZoneAmericaHermosillo
	TimeZoneAmericaIndianaIndianapolis
	TimeZoneAmericaIndianaKnox
	TimeZoneAmericaIndianaMarengo
	TimeZoneAmericaIndianaPetersburg
	TimeZoneAmericaIndianaTellCity
	TimeZoneAmericaIndianaVevay
	TimeZoneAmericaIndianaVincennes
	TimeZoneAmericaIndianaWinamac
	TimeZoneAmericaInuvik
	TimeZoneAmericaIqaluit
	TimeZoneAmericaJamaica
	TimeZoneAmericaJuneau
	TimeZoneAmericaKentuckyLouisville
	TimeZoneAmericaKentuckyMonticello
	TimeZoneAmericaKralendijk
	TimeZoneAmericaLaPaz
	TimeZoneAmericaLima
	TimeZoneAmericaLosAngeles
	TimeZoneAmericaLowerPrinces
	TimeZoneAmericaMaceio
	TimeZoneAmericaManagua
	TimeZoneAmericaManaus
	TimeZoneAmericaMarigot
	TimeZoneAmericaMartinique
	TimeZoneAmericaMatamoros
	TimeZoneAmericaMazatlan
	TimeZoneAmericaMenominee
	TimeZoneAmericaMerida
	TimeZoneAmericaMetlakatla
	TimeZoneAmericaMexicoCity
	TimeZoneAmericaMiquelon
	TimeZoneAmericaMoncton
	TimeZoneAmericaMonterrey
	TimeZoneAmericaMontevideo
	TimeZoneAmericaMontserrat
	TimeZoneAmericaNassau
	TimeZoneAmericaNewYork
	TimeZoneAmericaNipigon
	TimeZoneAmericaNome
	TimeZoneAmericaNoronha
	TimeZoneAmericaNorthDakotaBeulah
	TimeZoneAmericaNorthDakotaCenter
	TimeZoneAmericaNorthDakotaNewSalem
	TimeZoneAmericaNuuk
	TimeZoneAmericaOjinaga
	TimeZoneAmericaPanama
	TimeZoneAmericaPangnirtung
	TimeZoneAmericaParamaribo
	TimeZoneAmericaPhoenix
	TimeZoneAmericaPortauPrince
	TimeZoneAmericaPortofSpain
	TimeZoneAmericaPortoVelho
	TimeZoneAmericaPuertoRico
	TimeZoneAmericaPuntaArenas
	TimeZoneAmericaRainyRiver
	TimeZoneAmericaRankinInlet
	TimeZoneAmericaRecife
	TimeZoneAmericaRegina
	TimeZoneAmericaResolute
	TimeZoneAmericaRioBranco
	TimeZoneAmericaSantarem
	TimeZoneAmericaSantiago
	TimeZoneAmericaSantoDomingo
	TimeZoneAmericaSaoPaulo
	TimeZoneAmericaScoresbysund
	TimeZoneAmericaSitka
	TimeZoneAmericaStBarthelemy
	TimeZoneAmericaStJohns
	TimeZoneAmericaStKitts
	TimeZoneAmericaStLucia
	TimeZoneAmericaStThomas
	TimeZoneAmericaStVincent
	TimeZoneAmericaSwiftCurrent
	TimeZoneAmericaTegucigalpa
	TimeZoneAmericaThule
	TimeZoneAmericaThunderBay
	TimeZoneAmericaTijuana
	TimeZoneAmericaToronto
	TimeZoneAmericaTortola
	TimeZoneAmericaVancouver
	TimeZoneAmericaWhitehorse
	TimeZoneAmericaWinnipeg
	TimeZoneAmericaYakutat
	TimeZoneAmericaYellowknife
	TimeZoneAntarcticaCasey
	TimeZoneAntarcticaDavis
	TimeZoneAntarcticaDumontDUrville
	TimeZoneAntarcticaMacquarie
	TimeZoneAntarcticaMawson
	TimeZoneAntarcticaMcMurdo
	TimeZoneAntarcticaPalmer
	TimeZoneAntarcticaRothera
	TimeZoneAntarcticaSyowa
	TimeZoneAntarcticaTroll
	TimeZoneAntarcticaVostok
	TimeZoneArcticLongyearbyen
	TimeZoneAsiaAden
	TimeZoneAsiaAlmaty
	TimeZoneAsiaAmman
	TimeZoneAsiaAnadyr
	TimeZoneAsiaAqtau
	TimeZoneAsiaAqtobe
	TimeZoneAsiaAshgabat
	TimeZoneAsiaAtyrau
	TimeZoneAsiaBaghdad
	TimeZoneAsiaBahrain
	TimeZoneAsiaBaku
	TimeZoneAsiaBangkok
	TimeZoneAsiaBarnaul
	TimeZoneAsiaBeirut
	TimeZoneAsiaBishkek
	TimeZoneAsiaBrunei
	TimeZoneAsiaChita
	TimeZoneAsiaChoibalsan
	TimeZoneAsiaColombo
	TimeZoneAsiaDamascus
	TimeZoneAsiaDhaka
	TimeZoneAsiaDili
	TimeZoneAsiaDubai
	TimeZoneAsiaDushanbe
	TimeZoneAsiaFamagusta
	TimeZoneAsiaGaza
	TimeZoneAsiaHebron
	TimeZoneAsiaHoChiMinh
	TimeZoneAsiaHongKong
	TimeZoneAsiaHovd
	TimeZoneAsiaIrkutsk
	TimeZoneAsiaIstanbul
	TimeZoneAsiaJakarta
	TimeZoneAsiaJayapura
	TimeZoneAsiaJerusalem
	TimeZoneAsiaKabul
	TimeZoneAsiaKamchatka
	TimeZoneAsiaKarachi
	TimeZoneAsiaKathmandu
	TimeZoneAsiaKhandyga
	TimeZoneAsiaKolkata
	TimeZoneAsiaKrasnoyarsk
	TimeZoneAsiaKualaLumpur
	TimeZoneAsiaKuching
	TimeZoneAsiaKuwait
	TimeZoneAsiaMacau
	TimeZoneAsiaMagadan
	TimeZoneAsiaMakassar
	TimeZoneAsiaManila
	TimeZoneAsiaMuscat
	TimeZoneAsiaNicosia
	TimeZoneAsiaNovokuznetsk
	TimeZoneAsiaNovosibirsk
	TimeZoneAsiaOmsk
	TimeZoneAsiaOral
	TimeZoneAsiaPhnomPenh
	TimeZoneAsiaPontianak
	TimeZoneAsiaPyongyang
	TimeZoneAsiaQatar
	TimeZoneAsiaQostanay
	TimeZoneAsiaQyzylorda
	TimeZoneAsiaRangoon
	TimeZoneAsiaRiyadh
	TimeZoneAsiaSakhalin
	TimeZoneAsiaSamarkand
	TimeZoneAsiaSeoul
	TimeZoneAsiaShanghai
	TimeZoneAsiaSingapore
	TimeZoneAsiaSrednekolymsk
	TimeZoneAsiaTaipei
	TimeZoneAsiaTashkent
	TimeZoneAsiaTbilisi
	TimeZoneAsiaTehran
	TimeZoneAsiaThimphu
	TimeZoneAsiaTokyo
	TimeZoneAsiaTomsk
	TimeZoneAsiaUlaanbaatar
	TimeZoneAsiaUrumqi
	TimeZoneAsiaUstNera
	TimeZoneAsiaVientiane
	TimeZoneAsiaVladivostok
	TimeZoneAsiaYakutsk
	TimeZoneAsiaYangon
	TimeZoneAsiaYekaterinburg
	TimeZoneAsiaYerevan
	TimeZoneAtlanticAzores
	TimeZoneAtlanticBermuda
	TimeZoneAtlanticCanary
	TimeZoneAtlanticCapeVerde
	TimeZoneAtlanticFaroe
	TimeZoneAtlanticMadeira
	TimeZoneAtlanticReykjavik
	TimeZoneAtlanticSouthGeorgia
	TimeZoneAtlanticStHelena
	TimeZoneAtlanticStanley
	TimeZoneAustraliaAdelaide
	TimeZoneAustraliaBrisbane
	TimeZoneAustraliaBrokenHill
	TimeZoneAustraliaCurrie
	TimeZoneAustraliaDarwin
	TimeZoneAustraliaEucla
	TimeZoneAustraliaHobart
	TimeZoneAustraliaLindeman
	TimeZoneAustraliaLordHowe
	TimeZoneAustraliaMelbourne
	TimeZoneAustraliaPerth
	TimeZoneAustraliaSydney
	TimeZoneEuropeAmsterdam
	TimeZoneEuropeAndorra
	TimeZoneEuropeAstrakhan
	TimeZoneEuropeAthens
	TimeZoneEuropeBelgrade
	TimeZoneEuropeBerlin
	TimeZoneEuropeBratislava
	TimeZoneEuropeBrussels
	TimeZoneEuropeBucharest
	TimeZoneEuropeBudapest
	TimeZoneEuropeBusingen
	TimeZoneEuropeChisinau
	TimeZoneEuropeCopenhagen
	TimeZoneEuropeDublin
	TimeZoneEuropeGibraltar
	TimeZoneEuropeGuernsey
	TimeZoneEuropeHelsinki
	TimeZoneEuropeIsleofMan
	TimeZoneEuropeIstanbul
	TimeZoneEuropeJersey
	TimeZoneEuropeKaliningrad
	TimeZoneEuropeKiev
	TimeZoneEuropeKirov
	TimeZoneEuropeLisbon
	TimeZoneEuropeLjubljana
	TimeZoneEuropeLondon
	TimeZoneEuropeLuxembourg
	TimeZoneEuropeMadrid
	TimeZoneEuropeMalta
	TimeZoneEuropeMariehamn
	TimeZoneEuropeMinsk
	TimeZoneEuropeMonaco
	TimeZoneEuropeMoscow
	TimeZoneEuropeNicosia
	TimeZoneEuropeOslo
	TimeZoneEuropeParis
	TimeZoneEuropePodgorica
	TimeZoneEuropePrague
	TimeZoneEuropeRiga
	TimeZoneEuropeRome
	TimeZoneEuropeSamara
	TimeZoneEuropeSanMarino
	TimeZoneEuropeSarajevo
	TimeZoneEuropeSaratov
	TimeZoneEuropeSimferopol
	TimeZoneEuropeSkopje
	TimeZoneEuropeSofia
	TimeZoneEuropeStockholm
	TimeZoneEuropeTallinn
	TimeZoneEuropeTirane
	TimeZoneEuropeUlyanovsk
	TimeZoneEuropeUzhgorod
	TimeZoneEuropeVaduz
	TimeZoneEuropeVatican
	TimeZoneEuropeVienna
	TimeZoneEuropeVilnius
	TimeZoneEuropeVolgograd
	TimeZoneEuropeWarsaw
	TimeZoneEuropeZagreb
	TimeZoneEuropeZaporozhye
	TimeZoneEuropeZurich
	TimeZoneIndianAntananarivo
	TimeZoneIndianChagos
	TimeZoneIndianChristmas
	TimeZoneIndianCocos
	TimeZoneIndianComoro
	TimeZoneIndianKerguelen
	TimeZoneIndianMahe
	TimeZoneIndianMaldives
	TimeZoneIndianMauritius
	TimeZoneIndianMayotte
	TimeZoneIndianReunion
	TimeZonePacificApia
	TimeZonePacificAuckland
	TimeZonePacificBougainville
	TimeZonePacificChatham
	TimeZonePacificChuuk
	TimeZonePacificEaster
	TimeZonePacificEfate
	TimeZonePacificEnderbury
	TimeZonePacificFakaofo
	TimeZonePacificFiji
	TimeZonePacificFunafuti
	TimeZonePacificGalapagos
	TimeZonePacificGambier
	TimeZonePacificGuadalcanal
	TimeZonePacificGuam
	TimeZonePacificHonolulu
	TimeZonePacificKanton
	TimeZonePacificKiritimati
	TimeZonePacificKosrae
	TimeZonePacificKwajalein
	TimeZonePacificMajuro
	TimeZonePacificMarquesas
	TimeZonePacificMidway
	TimeZonePacificNauru
	TimeZonePacificNiue
	TimeZonePacificNorfolk
	TimeZonePacificNoumea
	TimeZonePacificPagoPago
	TimeZonePacificPalau
	TimeZonePacificPitcairn
	TimeZonePacificPohnpei
	TimeZonePacificPortMoresby
	TimeZonePacificRarotonga
	TimeZonePacificSaipan
	TimeZonePacificTahiti
	TimeZonePacificTarawa
	TimeZonePacificTongatapu
	TimeZonePacificWake
	TimeZonePacificWallis
)

func (tz TimeZone) String() string {
	switch tz {
	case TimeZoneUTC:
		return "utc"
	case TimeZoneClient:
		return "client"
	case TimeZoneAfricaAbidjan:
		return "Africa/Abidjan"
	case TimeZoneAfricaAccra:
		return "Africa/Accra"
	case TimeZoneAfricaAddisAbaba:
		return "Africa/Addis_Ababa"
	case TimeZoneAfricaAlgiers:
		return "Africa/Algiers"
	case TimeZoneAfricaAsmara:
		return "Africa/Asmara"
	case TimeZoneAfricaBamako:
		return "Africa/Bamako"
	case TimeZoneAfricaBangui:
		return "Africa/Bangui"
	case TimeZoneAfricaBanjul:
		return "Africa/Banjul"
	case TimeZoneAfricaBissau:
		return "Africa/Bissau"
	case TimeZoneAfricaBlantyre:
		return "Africa/Blantyre"
	case TimeZoneAfricaBrazzaville:
		return "Africa/Brazzaville"
	case TimeZoneAfricaBujumbura:
		return "Africa/Bujumbura"
	case TimeZoneAfricaCairo:
		return "Africa/Cairo"
	case TimeZoneAfricaCasablanca:
		return "Africa/Casablanca"
	case TimeZoneAfricaCeuta:
		return "Africa/Ceuta"
	case TimeZoneAfricaConakry:
		return "Africa/Conakry"
	case TimeZoneAfricaDakar:
		return "Africa/Dakar"
	case TimeZoneAfricaDaresSalaam:
		return "Africa/Dar_es_Salaam"
	case TimeZoneAfricaDjibouti:
		return "Africa/Djibouti"
	case TimeZoneAfricaDouala:
		return "Africa/Douala"
	case TimeZoneAfricaElAaiun:
		return "Africa/El_Aaiun"
	case TimeZoneAfricaFreetown:
		return "Africa/Freetown"
	case TimeZoneAfricaGaborone:
		return "Africa/Gaborone"
	case TimeZoneAfricaHarare:
		return "Africa/Harare"
	case TimeZoneAfricaJohannesburg:
		return "Africa/Johannesburg"
	case TimeZoneAfricaJuba:
		return "Africa/Juba"
	case TimeZoneAfricaKampala:
		return "Africa/Kampala"
	case TimeZoneAfricaKhartoum:
		return "Africa/Khartoum"
	case TimeZoneAfricaKigali:
		return "Africa/Kigali"
	case TimeZoneAfricaKinshasa:
		return "Africa/Kinshasa"
	case TimeZoneAfricaLagos:
		return "Africa/Lagos"
	case TimeZoneAfricaLibreville:
		return "Africa/Libreville"
	case TimeZoneAfricaLome:
		return "Africa/Lome"
	case TimeZoneAfricaLuanda:
		return "Africa/Luanda"
	case TimeZoneAfricaLubumbashi:
		return "Africa/Lubumbashi"
	case TimeZoneAfricaLusaka:
		return "Africa/Lusaka"
	case TimeZoneAfricaMalabo:
		return "Africa/Malabo"
	case TimeZoneAfricaMaputo:
		return "Africa/Maputo"
	case TimeZoneAfricaMaseru:
		return "Africa/Maseru"
	case TimeZoneAfricaMbabane:
		return "Africa/Mbabane"
	case TimeZoneAfricaMogadishu:
		return "Africa/Mogadishu"
	case TimeZoneAfricaMonrovia:
		return "Africa/Monrovia"
	case TimeZoneAfricaNairobi:
		return "Africa/Nairobi"
	case TimeZoneAfricaNdjamena:
		return "Africa/Ndjamena"
	case TimeZoneAfricaNiamey:
		return "Africa/Niamey"
	case TimeZoneAfricaNouakchott:
		return "Africa/Nouakchott"
	case TimeZoneAfricaOuagadougou:
		return "Africa/Ouagadougou"
	case TimeZoneAfricaPortoNovo:
		return "Africa/Porto-Novo"
	case TimeZoneAfricaSaoTome:
		return "Africa/Sao_Tome"
	case TimeZoneAfricaTripoli:
		return "Africa/Tripoli"
	case TimeZoneAfricaTunis:
		return "Africa/Tunis"
	case TimeZoneAfricaWindhoek:
		return "Africa/Windhoek"
	case TimeZoneAmericaAdak:
		return "America/Adak"
	case TimeZoneAmericaAnchorage:
		return "America/Anchorage"
	case TimeZoneAmericaAnguilla:
		return "America/Anguilla"
	case TimeZoneAmericaAntigua:
		return "America/Antigua"
	case TimeZoneAmericaAraguaina:
		return "America/Araguaina"
	case TimeZoneAmericaArgentinaBuenosAires:
		return "America/Argentina/Buenos_Aires"
	case TimeZoneAmericaArgentinaCatamarca:
		return "America/Argentina/Catamarca"
	case TimeZoneAmericaArgentinaCordoba:
		return "America/Argentina/Cordoba"
	case TimeZoneAmericaArgentinaJujuy:
		return "America/Argentina/Jujuy"
	case TimeZoneAmericaArgentinaLaRioja:
		return "America/Argentina/La_Rioja"
	case TimeZoneAmericaArgentinaMendoza:
		return "America/Argentina/Mendoza"
	case TimeZoneAmericaArgentinaRioGallegos:
		return "America/Argentina/Rio_Gallegos"
	case TimeZoneAmericaArgentinaSalta:
		return "America/Argentina/Salta"
	case TimeZoneAmericaArgentinaSanJuan:
		return "America/Argentina/San_Juan"
	case TimeZoneAmericaArgentinaSanLuis:
		return "America/Argentina/San_Luis"
	case TimeZoneAmericaArgentinaTucuman:
		return "America/Argentina/Tucuman"
	case TimeZoneAmericaArgentinaUshuaia:
		return "America/Argentina/Ushuaia"
	case TimeZoneAmericaAruba:
		return "America/Aruba"
	case TimeZoneAmericaAsuncion:
		return "America/Asuncion"
	case TimeZoneAmericaAtikokan:
		return "America/Atikokan"
	case TimeZoneAmericaBahia:
		return "America/Bahia"
	case TimeZoneAmericaBahiaBanderas:
		return "America/Bahia_Banderas"
	case TimeZoneAmericaBarbados:
		return "America/Barbados"
	case TimeZoneAmericaBelem:
		return "America/Belem"
	case TimeZoneAmericaBelize:
		return "America/Belize"
	case TimeZoneAmericaBlancSablon:
		return "America/Blanc-Sablon"
	case TimeZoneAmericaBoaVista:
		return "America/Boa_Vista"
	case TimeZoneAmericaBogota:
		return "America/Bogota"
	case TimeZoneAmericaBoise:
		return "America/Boise"
	case TimeZoneAmericaCambridgeBay:
		return "America/Cambridge_Bay"
	case TimeZoneAmericaCampoGrande:
		return "America/Campo_Grande"
	case TimeZoneAmericaCancun:
		return "America/Cancun"
	case TimeZoneAmericaCaracas:
		return "America/Caracas"
	case TimeZoneAmericaCayenne:
		return "America/Cayenne"
	case TimeZoneAmericaCayman:
		return "America/Cayman"
	case TimeZoneAmericaChicago:
		return "America/Chicago"
	case TimeZoneAmericaChihuahua:
		return "America/Chihuahua"
	case TimeZoneAmericaCostaRica:
		return "America/Costa_Rica"
	case TimeZoneAmericaCreston:
		return "America/Creston"
	case TimeZoneAmericaCuiaba:
		return "America/Cuiaba"
	case TimeZoneAmericaCuracao:
		return "America/Curacao"
	case TimeZoneAmericaDanmarkshavn:
		return "America/Danmarkshavn"
	case TimeZoneAmericaDawson:
		return "America/Dawson"
	case TimeZoneAmericaDawsonCreek:
		return "America/Dawson_Creek"
	case TimeZoneAmericaDenver:
		return "America/Denver"
	case TimeZoneAmericaDetroit:
		return "America/Detroit"
	case TimeZoneAmericaDominica:
		return "America/Dominica"
	case TimeZoneAmericaEdmonton:
		return "America/Edmonton"
	case TimeZoneAmericaEirunepe:
		return "America/Eirunepe"
	case TimeZoneAmericaElSalvador:
		return "America/El_Salvador"
	case TimeZoneAmericaFortNelson:
		return "America/Fort_Nelson"
	case TimeZoneAmericaFortaleza:
		return "America/Fortaleza"
	case TimeZoneAmericaGlaceBay:
		return "America/Glace_Bay"
	case TimeZoneAmericaGodthab:
		return "America/Godthab"
	case TimeZoneAmericaGooseBay:
		return "America/Goose_Bay"
	case TimeZoneAmericaGrandTurk:
		return "America/Grand_Turk"
	case TimeZoneAmericaGrenada:
		return "America/Grenada"
	case TimeZoneAmericaGuadeloupe:
		return "America/Guadeloupe"
	case TimeZoneAmericaGuatemala:
		return "America/Guatemala"
	case TimeZoneAmericaGuayaquil:
		return "America/Guayaquil"
	case TimeZoneAmericaGuyana:
		return "America/Guyana"
	case TimeZoneAmericaHalifax:
		return "America/Halifax"
	case TimeZoneAmericaHavana:
		return "America/Havana"
	case TimeZoneAmericaHermosillo:
		return "America/Hermosillo"
	case TimeZoneAmericaIndianaIndianapolis:
		return "America/Indiana/Indianapolis"
	case TimeZoneAmericaIndianaKnox:
		return "America/Indiana/Knox"
	case TimeZoneAmericaIndianaMarengo:
		return "America/Indiana/Marengo"
	case TimeZoneAmericaIndianaPetersburg:
		return "America/Indiana/Petersburg"
	case TimeZoneAmericaIndianaTellCity:
		return "America/Indiana/Tell_City"
	case TimeZoneAmericaIndianaVevay:
		return "America/Indiana/Vevay"
	case TimeZoneAmericaIndianaVincennes:
		return "America/Indiana/Vincennes"
	case TimeZoneAmericaIndianaWinamac:
		return "America/Indiana/Winamac"
	case TimeZoneAmericaInuvik:
		return "America/Inuvik"
	case TimeZoneAmericaIqaluit:
		return "America/Iqaluit"
	case TimeZoneAmericaJamaica:
		return "America/Jamaica"
	case TimeZoneAmericaJuneau:
		return "America/Juneau"
	case TimeZoneAmericaKentuckyLouisville:
		return "America/Kentucky/Louisville"
	case TimeZoneAmericaKentuckyMonticello:
		return "America/Kentucky/Monticello"
	case TimeZoneAmericaKralendijk:
		return "America/Kralendijk"
	case TimeZoneAmericaLaPaz:
		return "America/La_Paz"
	case TimeZoneAmericaLima:
		return "America/Lima"
	case TimeZoneAmericaLosAngeles:
		return "America/Los_Angeles"
	case TimeZoneAmericaLowerPrinces:
		return "America/Lower_Princes"
	case TimeZoneAmericaMaceio:
		return "America/Maceio"
	case TimeZoneAmericaManagua:
		return "America/Managua"
	case TimeZoneAmericaManaus:
		return "America/Manaus"
	case TimeZoneAmericaMarigot:
		return "America/Marigot"
	case TimeZoneAmericaMartinique:
		return "America/Martinique"
	case TimeZoneAmericaMatamoros:
		return "America/Matamoros"
	case TimeZoneAmericaMazatlan:
		return "America/Mazatlan"
	case TimeZoneAmericaMenominee:
		return "America/Menominee"
	case TimeZoneAmericaMerida:
		return "America/Merida"
	case TimeZoneAmericaMetlakatla:
		return "America/Metlakatla"
	case TimeZoneAmericaMexicoCity:
		return "America/Mexico_City"
	case TimeZoneAmericaMiquelon:
		return "America/Miquelon"
	case TimeZoneAmericaMoncton:
		return "America/Moncton"
	case TimeZoneAmericaMonterrey:
		return "America/Monterrey"
	case TimeZoneAmericaMontevideo:
		return "America/Montevideo"
	case TimeZoneAmericaMontserrat:
		return "America/Montserrat"
	case TimeZoneAmericaNassau:
		return "America/Nassau"
	case TimeZoneAmericaNewYork:
		return "America/New_York"
	case TimeZoneAmericaNipigon:
		return "America/Nipigon"
	case TimeZoneAmericaNome:
		return "America/Nome"
	case TimeZoneAmericaNoronha:
		return "America/Noronha"
	case TimeZoneAmericaNorthDakotaBeulah:
		return "America/North_Dakota/Beulah"
	case TimeZoneAmericaNorthDakotaCenter:
		return "America/North_Dakota/Center"
	case TimeZoneAmericaNorthDakotaNewSalem:
		return "America/North_Dakota/New_Salem"
	case TimeZoneAmericaNuuk:
		return "America/Nuuk"
	case TimeZoneAmericaOjinaga:
		return "America/Ojinaga"
	case TimeZoneAmericaPanama:
		return "America/Panama"
	case TimeZoneAmericaPangnirtung:
		return "America/Pangnirtung"
	case TimeZoneAmericaParamaribo:
		return "America/Paramaribo"
	case TimeZoneAmericaPhoenix:
		return "America/Phoenix"
	case TimeZoneAmericaPortauPrince:
		return "America/Port-au-Prince"
	case TimeZoneAmericaPortofSpain:
		return "America/Port_of_Spain"
	case TimeZoneAmericaPortoVelho:
		return "America/Porto_Velho"
	case TimeZoneAmericaPuertoRico:
		return "America/Puerto_Rico"
	case TimeZoneAmericaPuntaArenas:
		return "America/Punta_Arenas"
	case TimeZoneAmericaRainyRiver:
		return "America/Rainy_River"
	case TimeZoneAmericaRankinInlet:
		return "America/Rankin_Inlet"
	case TimeZoneAmericaRecife:
		return "America/Recife"
	case TimeZoneAmericaRegina:
		return "America/Regina"
	case TimeZoneAmericaResolute:
		return "America/Resolute"
	case TimeZoneAmericaRioBranco:
		return "America/Rio_Branco"
	case TimeZoneAmericaSantarem:
		return "America/Santarem"
	case TimeZoneAmericaSantiago:
		return "America/Santiago"
	case TimeZoneAmericaSantoDomingo:
		return "America/Santo_Domingo"
	case TimeZoneAmericaSaoPaulo:
		return "America/Sao_Paulo"
	case TimeZoneAmericaScoresbysund:
		return "America/Scoresbysund"
	case TimeZoneAmericaSitka:
		return "America/Sitka"
	case TimeZoneAmericaStBarthelemy:
		return "America/St_Barthelemy"
	case TimeZoneAmericaStJohns:
		return "America/St_Johns"
	case TimeZoneAmericaStKitts:
		return "America/St_Kitts"
	case TimeZoneAmericaStLucia:
		return "America/St_Lucia"
	case TimeZoneAmericaStThomas:
		return "America/St_Thomas"
	case TimeZoneAmericaStVincent:
		return "America/St_Vincent"
	case TimeZoneAmericaSwiftCurrent:
		return "America/Swift_Current"
	case TimeZoneAmericaTegucigalpa:
		return "America/Tegucigalpa"
	case TimeZoneAmericaThule:
		return "America/Thule"
	case TimeZoneAmericaThunderBay:
		return "America/Thunder_Bay"
	case TimeZoneAmericaTijuana:
		return "America/Tijuana"
	case TimeZoneAmericaToronto:
		return "America/Toronto"
	case TimeZoneAmericaTortola:
		return "America/Tortola"
	case TimeZoneAmericaVancouver:
		return "America/Vancouver"
	case TimeZoneAmericaWhitehorse:
		return "America/Whitehorse"
	case TimeZoneAmericaWinnipeg:
		return "America/Winnipeg"
	case TimeZoneAmericaYakutat:
		return "America/Yakutat"
	case TimeZoneAmericaYellowknife:
		return "America/Yellowknife"
	case TimeZoneAntarcticaCasey:
		return "Antarctica/Casey"
	case TimeZoneAntarcticaDavis:
		return "Antarctica/Davis"
	case TimeZoneAntarcticaDumontDUrville:
		return "Antarctica/DumontDUrville"
	case TimeZoneAntarcticaMacquarie:
		return "Antarctica/Macquarie"
	case TimeZoneAntarcticaMawson:
		return "Antarctica/Mawson"
	case TimeZoneAntarcticaMcMurdo:
		return "Antarctica/McMurdo"
	case TimeZoneAntarcticaPalmer:
		return "Antarctica/Palmer"
	case TimeZoneAntarcticaRothera:
		return "Antarctica/Rothera"
	case TimeZoneAntarcticaSyowa:
		return "Antarctica/Syowa"
	case TimeZoneAntarcticaTroll:
		return "Antarctica/Troll"
	case TimeZoneAntarcticaVostok:
		return "Antarctica/Vostok"
	case TimeZoneArcticLongyearbyen:
		return "Arctic/Longyearbyen"
	case TimeZoneAsiaAden:
		return "Asia/Aden"
	case TimeZoneAsiaAlmaty:
		return "Asia/Almaty"
	case TimeZoneAsiaAmman:
		return "Asia/Amman"
	case TimeZoneAsiaAnadyr:
		return "Asia/Anadyr"
	case TimeZoneAsiaAqtau:
		return "Asia/Aqtau"
	case TimeZoneAsiaAqtobe:
		return "Asia/Aqtobe"
	case TimeZoneAsiaAshgabat:
		return "Asia/Ashgabat"
	case TimeZoneAsiaAtyrau:
		return "Asia/Atyrau"
	case TimeZoneAsiaBaghdad:
		return "Asia/Baghdad"
	case TimeZoneAsiaBahrain:
		return "Asia/Bahrain"
	case TimeZoneAsiaBaku:
		return "Asia/Baku"
	case TimeZoneAsiaBangkok:
		return "Asia/Bangkok"
	case TimeZoneAsiaBarnaul:
		return "Asia/Barnaul"
	case TimeZoneAsiaBeirut:
		return "Asia/Beirut"
	case TimeZoneAsiaBishkek:
		return "Asia/Bishkek"
	case TimeZoneAsiaBrunei:
		return "Asia/Brunei"
	case TimeZoneAsiaChita:
		return "Asia/Chita"
	case TimeZoneAsiaChoibalsan:
		return "Asia/Choibalsan"
	case TimeZoneAsiaColombo:
		return "Asia/Colombo"
	case TimeZoneAsiaDamascus:
		return "Asia/Damascus"
	case TimeZoneAsiaDhaka:
		return "Asia/Dhaka"
	case TimeZoneAsiaDili:
		return "Asia/Dili"
	case TimeZoneAsiaDubai:
		return "Asia/Dubai"
	case TimeZoneAsiaDushanbe:
		return "Asia/Dushanbe"
	case TimeZoneAsiaFamagusta:
		return "Asia/Famagusta"
	case TimeZoneAsiaGaza:
		return "Asia/Gaza"
	case TimeZoneAsiaHebron:
		return "Asia/Hebron"
	case TimeZoneAsiaHoChiMinh:
		return "Asia/Ho_Chi_Minh"
	case TimeZoneAsiaHongKong:
		return "Asia/Hong_Kong"
	case TimeZoneAsiaHovd:
		return "Asia/Hovd"
	case TimeZoneAsiaIrkutsk:
		return "Asia/Irkutsk"
	case TimeZoneAsiaIstanbul:
		return "Asia/Istanbul"
	case TimeZoneAsiaJakarta:
		return "Asia/Jakarta"
	case TimeZoneAsiaJayapura:
		return "Asia/Jayapura"
	case TimeZoneAsiaJerusalem:
		return "Asia/Jerusalem"
	case TimeZoneAsiaKabul:
		return "Asia/Kabul"
	case TimeZoneAsiaKamchatka:
		return "Asia/Kamchatka"
	case TimeZoneAsiaKarachi:
		return "Asia/Karachi"
	case TimeZoneAsiaKathmandu:
		return "Asia/Kathmandu"
	case TimeZoneAsiaKhandyga:
		return "Asia/Khandyga"
	case TimeZoneAsiaKolkata:
		return "Asia/Kolkata"
	case TimeZoneAsiaKrasnoyarsk:
		return "Asia/Krasnoyarsk"
	case TimeZoneAsiaKualaLumpur:
		return "Asia/Kuala_Lumpur"
	case TimeZoneAsiaKuching:
		return "Asia/Kuching"
	case TimeZoneAsiaKuwait:
		return "Asia/Kuwait"
	case TimeZoneAsiaMacau:
		return "Asia/Macau"
	case TimeZoneAsiaMagadan:
		return "Asia/Magadan"
	case TimeZoneAsiaMakassar:
		return "Asia/Makassar"
	case TimeZoneAsiaManila:
		return "Asia/Manila"
	case TimeZoneAsiaMuscat:
		return "Asia/Muscat"
	case TimeZoneAsiaNicosia:
		return "Asia/Nicosia"
	case TimeZoneAsiaNovokuznetsk:
		return "Asia/Novokuznetsk"
	case TimeZoneAsiaNovosibirsk:
		return "Asia/Novosibirsk"
	case TimeZoneAsiaOmsk:
		return "Asia/Omsk"
	case TimeZoneAsiaOral:
		return "Asia/Oral"
	case TimeZoneAsiaPhnomPenh:
		return "Asia/Phnom_Penh"
	case TimeZoneAsiaPontianak:
		return "Asia/Pontianak"
	case TimeZoneAsiaPyongyang:
		return "Asia/Pyongyang"
	case TimeZoneAsiaQatar:
		return "Asia/Qatar"
	case TimeZoneAsiaQostanay:
		return "Asia/Qostanay"
	case TimeZoneAsiaQyzylorda:
		return "Asia/Qyzylorda"
	case TimeZoneAsiaRangoon:
		return "Asia/Rangoon"
	case TimeZoneAsiaRiyadh:
		return "Asia/Riyadh"
	case TimeZoneAsiaSakhalin:
		return "Asia/Sakhalin"
	case TimeZoneAsiaSamarkand:
		return "Asia/Samarkand"
	case TimeZoneAsiaSeoul:
		return "Asia/Seoul"
	case TimeZoneAsiaShanghai:
		return "Asia/Shanghai"
	case TimeZoneAsiaSingapore:
		return "Asia/Singapore"
	case TimeZoneAsiaSrednekolymsk:
		return "Asia/Srednekolymsk"
	case TimeZoneAsiaTaipei:
		return "Asia/Taipei"
	case TimeZoneAsiaTashkent:
		return "Asia/Tashkent"
	case TimeZoneAsiaTbilisi:
		return "Asia/Tbilisi"
	case TimeZoneAsiaTehran:
		return "Asia/Tehran"
	case TimeZoneAsiaThimphu:
		return "Asia/Thimphu"
	case TimeZoneAsiaTokyo:
		return "Asia/Tokyo"
	case TimeZoneAsiaTomsk:
		return "Asia/Tomsk"
	case TimeZoneAsiaUlaanbaatar:
		return "Asia/Ulaanbaatar"
	case TimeZoneAsiaUrumqi:
		return "Asia/Urumqi"
	case TimeZoneAsiaUstNera:
		return "Asia/Ust-Nera"
	case TimeZoneAsiaVientiane:
		return "Asia/Vientiane"
	case TimeZoneAsiaVladivostok:
		return "Asia/Vladivostok"
	case TimeZoneAsiaYakutsk:
		return "Asia/Yakutsk"
	case TimeZoneAsiaYangon:
		return "Asia/Yangon"
	case TimeZoneAsiaYekaterinburg:
		return "Asia/Yekaterinburg"
	case TimeZoneAsiaYerevan:
		return "Asia/Yerevan"
	case TimeZoneAtlanticAzores:
		return "Atlantic/Azores"
	case TimeZoneAtlanticBermuda:
		return "Atlantic/Bermuda"
	case TimeZoneAtlanticCanary:
		return "Atlantic/Canary"
	case TimeZoneAtlanticCapeVerde:
		return "Atlantic/Cape_Verde"
	case TimeZoneAtlanticFaroe:
		return "Atlantic/Faroe"
	case TimeZoneAtlanticMadeira:
		return "Atlantic/Madeira"
	case TimeZoneAtlanticReykjavik:
		return "Atlantic/Reykjavik"
	case TimeZoneAtlanticSouthGeorgia:
		return "Atlantic/South_Georgia"
	case TimeZoneAtlanticStHelena:
		return "Atlantic/St_Helena"
	case TimeZoneAtlanticStanley:
		return "Atlantic/Stanley"
	case TimeZoneAustraliaAdelaide:
		return "Australia/Adelaide"
	case TimeZoneAustraliaBrisbane:
		return "Australia/Brisbane"
	case TimeZoneAustraliaBrokenHill:
		return "Australia/Broken_Hill"
	case TimeZoneAustraliaCurrie:
		return "Australia/Currie"
	case TimeZoneAustraliaDarwin:
		return "Australia/Darwin"
	case TimeZoneAustraliaEucla:
		return "Australia/Eucla"
	case TimeZoneAustraliaHobart:
		return "Australia/Hobart"
	case TimeZoneAustraliaLindeman:
		return "Australia/Lindeman"
	case TimeZoneAustraliaLordHowe:
		return "Australia/Lord_Howe"
	case TimeZoneAustraliaMelbourne:
		return "Australia/Melbourne"
	case TimeZoneAustraliaPerth:
		return "Australia/Perth"
	case TimeZoneAustraliaSydney:
		return "Australia/Sydney"
	case TimeZoneEuropeAmsterdam:
		return "Europe/Amsterdam"
	case TimeZoneEuropeAndorra:
		return "Europe/Andorra"
	case TimeZoneEuropeAstrakhan:
		return "Europe/Astrakhan"
	case TimeZoneEuropeAthens:
		return "Europe/Athens"
	case TimeZoneEuropeBelgrade:
		return "Europe/Belgrade"
	case TimeZoneEuropeBerlin:
		return "Europe/Berlin"
	case TimeZoneEuropeBratislava:
		return "Europe/Bratislava"
	case TimeZoneEuropeBrussels:
		return "Europe/Brussels"
	case TimeZoneEuropeBucharest:
		return "Europe/Bucharest"
	case TimeZoneEuropeBudapest:
		return "Europe/Budapest"
	case TimeZoneEuropeBusingen:
		return "Europe/Busingen"
	case TimeZoneEuropeChisinau:
		return "Europe/Chisinau"
	case TimeZoneEuropeCopenhagen:
		return "Europe/Copenhagen"
	case TimeZoneEuropeDublin:
		return "Europe/Dublin"
	case TimeZoneEuropeGibraltar:
		return "Europe/Gibraltar"
	case TimeZoneEuropeGuernsey:
		return "Europe/Guernsey"
	case TimeZoneEuropeHelsinki:
		return "Europe/Helsinki"
	case TimeZoneEuropeIsleofMan:
		return "Europe/Isle_of_Man"
	case TimeZoneEuropeIstanbul:
		return "Europe/Istanbul"
	case TimeZoneEuropeJersey:
		return "Europe/Jersey"
	case TimeZoneEuropeKaliningrad:
		return "Europe/Kaliningrad"
	case TimeZoneEuropeKiev:
		return "Europe/Kiev"
	case TimeZoneEuropeKirov:
		return "Europe/Kirov"
	case TimeZoneEuropeLisbon:
		return "Europe/Lisbon"
	case TimeZoneEuropeLjubljana:
		return "Europe/Ljubljana"
	case TimeZoneEuropeLondon:
		return "Europe/London"
	case TimeZoneEuropeLuxembourg:
		return "Europe/Luxembourg"
	case TimeZoneEuropeMadrid:
		return "Europe/Madrid"
	case TimeZoneEuropeMalta:
		return "Europe/Malta"
	case TimeZoneEuropeMariehamn:
		return "Europe/Mariehamn"
	case TimeZoneEuropeMinsk:
		return "Europe/Minsk"
	case TimeZoneEuropeMonaco:
		return "Europe/Monaco"
	case TimeZoneEuropeMoscow:
		return "Europe/Moscow"
	case TimeZoneEuropeNicosia:
		return "Europe/Nicosia"
	case TimeZoneEuropeOslo:
		return "Europe/Oslo"
	case TimeZoneEuropeParis:
		return "Europe/Paris"
	case TimeZoneEuropePodgorica:
		return "Europe/Podgorica"
	case TimeZoneEuropePrague:
		return "Europe/Prague"
	case TimeZoneEuropeRiga:
		return "Europe/Riga"
	case TimeZoneEuropeRome:
		return "Europe/Rome"
	case TimeZoneEuropeSamara:
		return "Europe/Samara"
	case TimeZoneEuropeSanMarino:
		return "Europe/San_Marino"
	case TimeZoneEuropeSarajevo:
		return "Europe/Sarajevo"
	case TimeZoneEuropeSaratov:
		return "Europe/Saratov"
	case TimeZoneEuropeSimferopol:
		return "Europe/Simferopol"
	case TimeZoneEuropeSkopje:
		return "Europe/Skopje"
	case TimeZoneEuropeSofia:
		return "Europe/Sofia"
	case TimeZoneEuropeStockholm:
		return "Europe/Stockholm"
	case TimeZoneEuropeTallinn:
		return "Europe/Tallinn"
	case TimeZoneEuropeTirane:
		return "Europe/Tirane"
	case TimeZoneEuropeUlyanovsk:
		return "Europe/Ulyanovsk"
	case TimeZoneEuropeUzhgorod:
		return "Europe/Uzhgorod"
	case TimeZoneEuropeVaduz:
		return "Europe/Vaduz"
	case TimeZoneEuropeVatican:
		return "Europe/Vatican"
	case TimeZoneEuropeVienna:
		return "Europe/Vienna"
	case TimeZoneEuropeVilnius:
		return "Europe/Vilnius"
	case TimeZoneEuropeVolgograd:
		return "Europe/Volgograd"
	case TimeZoneEuropeWarsaw:
		return "Europe/Warsaw"
	case TimeZoneEuropeZagreb:
		return "Europe/Zagreb"
	case TimeZoneEuropeZaporozhye:
		return "Europe/Zaporozhye"
	case TimeZoneEuropeZurich:
		return "Europe/Zurich"
	case TimeZoneIndianAntananarivo:
		return "Indian/Antananarivo"
	case TimeZoneIndianChagos:
		return "Indian/Chagos"
	case TimeZoneIndianChristmas:
		return "Indian/Christmas"
	case TimeZoneIndianCocos:
		return "Indian/Cocos"
	case TimeZoneIndianComoro:
		return "Indian/Comoro"
	case TimeZoneIndianKerguelen:
		return "Indian/Kerguelen"
	case TimeZoneIndianMahe:
		return "Indian/Mahe"
	case TimeZoneIndianMaldives:
		return "Indian/Maldives"
	case TimeZoneIndianMauritius:
		return "Indian/Mauritius"
	case TimeZoneIndianMayotte:
		return "Indian/Mayotte"
	case TimeZoneIndianReunion:
		return "Indian/Reunion"
	case TimeZonePacificApia:
		return "Pacific/Apia"
	case TimeZonePacificAuckland:
		return "Pacific/Auckland"
	case TimeZonePacificBougainville:
		return "Pacific/Bougainville"
	case TimeZonePacificChatham:
		return "Pacific/Chatham"
	case TimeZonePacificChuuk:
		return "Pacific/Chuuk"
	case TimeZonePacificEaster:
		return "Pacific/Easter"
	case TimeZonePacificEfate:
		return "Pacific/Efate"
	case TimeZonePacificEnderbury:
		return "Pacific/Enderbury"
	case TimeZonePacificFakaofo:
		return "Pacific/Fakaofo"
	case TimeZonePacificFiji:
		return "Pacific/Fiji"
	case TimeZonePacificFunafuti:
		return "Pacific/Funafuti"
	case TimeZonePacificGalapagos:
		return "Pacific/Galapagos"
	case TimeZonePacificGambier:
		return "Pacific/Gambier"
	case TimeZonePacificGuadalcanal:
		return "Pacific/Guadalcanal"
	case TimeZonePacificGuam:
		return "Pacific/Guam"
	case TimeZonePacificHonolulu:
		return "Pacific/Honolulu"
	case TimeZonePacificKanton:
		return "Pacific/Kanton"
	case TimeZonePacificKiritimati:
		return "Pacific/Kiritimati"
	case TimeZonePacificKosrae:
		return "Pacific/Kosrae"
	case TimeZonePacificKwajalein:
		return "Pacific/Kwajalein"
	case TimeZonePacificMajuro:
		return "Pacific/Majuro"
	case TimeZonePacificMarquesas:
		return "Pacific/Marquesas"
	case TimeZonePacificMidway:
		return "Pacific/Midway"
	case TimeZonePacificNauru:
		return "Pacific/Nauru"
	case TimeZonePacificNiue:
		return "Pacific/Niue"
	case TimeZonePacificNorfolk:
		return "Pacific/Norfolk"
	case TimeZonePacificNoumea:
		return "Pacific/Noumea"
	case TimeZonePacificPagoPago:
		return "Pacific/Pago_Pago"
	case TimeZonePacificPalau:
		return "Pacific/Palau"
	case TimeZonePacificPitcairn:
		return "Pacific/Pitcairn"
	case TimeZonePacificPohnpei:
		return "Pacific/Pohnpei"
	case TimeZonePacificPortMoresby:
		return "Pacific/Port_Moresby"
	case TimeZonePacificRarotonga:
		return "Pacific/Rarotonga"
	case TimeZonePacificSaipan:
		return "Pacific/Saipan"
	case TimeZonePacificTahiti:
		return "Pacific/Tahiti"
	case TimeZonePacificTarawa:
		return "Pacific/Tarawa"
	case TimeZonePacificTongatapu:
		return "Pacific/Tongatapu"
	case TimeZonePacificWake:
		return "Pacific/Wake"
	case TimeZonePacificWallis:
		return "Pacific/Wallis"
	}

	panic("Unrecognized time zone")
}

func (tz TimeZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(tz.String())
}
