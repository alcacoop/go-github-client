package client

type JsonData map[string]interface{}

func (j JsonData) GetString(attr string) string {
	return j[attr].(string)
}
