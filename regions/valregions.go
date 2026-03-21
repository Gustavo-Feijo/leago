package regions

// Valorant APIs have different routings.
type ValRegion string

const (
	ValRegionAP      ValRegion = "ap"
	ValRegionBR      ValRegion = "br"
	ValRegionEsports ValRegion = "esports"
	ValRegionEU      ValRegion = "eu"
	ValRegionKR      ValRegion = "kr"
	ValRegionLatam   ValRegion = "latam"
	ValRegionNA      ValRegion = "na"
)
