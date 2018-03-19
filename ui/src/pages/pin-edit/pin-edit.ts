import { Component } from '@angular/core';
import { NavParams, ViewController, Events } from 'ionic-angular';
import { MapService } from '../../providers/services/map-service';

@Component({
  selector: 'pin-edit',
  templateUrl: 'pin-edit.html'
})

export class PinEdit {
  lat: number;
  lng: number;
  title: string;
  notes: string;
  mapId: number;

  constructor(public params: NavParams, public viewCtrl: ViewController, public mapSvc: MapService, public events: Events) {
    this.lat = params.get("lat");
    this.lng = params.get("lng");
    this.title = params.get("title");
    this.notes = params.get("notes");
    this.mapId = params.get("mapid");
  }

  dismiss(){
    this.viewCtrl.dismiss();
  }

  save(){
    console.log("lat:", this.lat, "lng:", this.lng, "title:", this.title, "notes:", this.notes, "mapid:", this.mapId);
    this.mapSvc.addPin(this.lat, this.lng, this.title, this.notes, this.mapId).then(point => {
      this.events.publish("interests:new", point)
    });
    this.viewCtrl.dismiss();
  }
}
