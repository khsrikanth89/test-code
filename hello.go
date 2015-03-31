package guestbook

import (
        "net/http"
        "appengine"
        "appengine/datastore"
        "encoding/json"
        "appengine/channel"
)


type Channelstore struct {
    ChannelId  string
    VidaoId string
}


func init() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){ http.ServeFile(w,r,"test.html")})
  http.HandleFunc("/createChannel", createChannel)
  http.HandleFunc("/message", messageHandler)
}

func getKey(c appengine.Context, VidaoId string) *datastore.Key {
  return datastore.NewKey(c, "channel", VidaoId, 0, nil)
}


func createChannel(w http.ResponseWriter, r *http.Request) {
  VidaoId := r.FormValue("user")
  c := appengine.NewContext(r)
  c.Infof("VidaoId",VidaoId)
  tok, err := channel.Create(c, VidaoId)
  c.Infof("Token",tok)
    if err != nil {
        http.Error(w, "Couldn't create Channel", http.StatusInternalServerError)
        c.Errorf("channel.Create: %v", err)
        return
    }
    g := Channelstore{
            ChannelId: tok,
            VidaoId: VidaoId,
    }
    _, err = datastore.Put(c, getKey(c,VidaoId), &g)
    if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
    }
  data,_ := json.Marshal(tok)
  w.Write(data)
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  // from := r.FormValue("from")
  to := r.FormValue("to")
  var g Channelstore
  if err := datastore.Get(c, getKey(c,to), &g); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }
  c.Infof("Channelstore",g);
  message := r.FormValue("message")
  channel.Send(c,g.VidaoId,message)
}

