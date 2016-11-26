package design

type Design struct {
	Custom map[string]DesignData
	Dsl    map[string]DesignData
	Type   map[string]DesignData
}

type DesignData map[interface{}]interface{}
