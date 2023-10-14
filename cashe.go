package main

import "time"

type CashedItem struct {
	Data interface{}
	TimeCashed time.Time 
}

type CasheStore struct {
	Items map[string]CashedItem
}

func (c *CasheStore) SetCashe(key string, data interface{}){
	c.Items[key] = CashedItem{Data: data, TimeCashed: time.Now()}
}

func (c *CasheStore) GetCashe(key string) (interface{}, bool) {

}