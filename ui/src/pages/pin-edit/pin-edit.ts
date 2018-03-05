import { Component } from '@angular/core';
import { NavParams, ViewController } from 'ionic-angular';
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

  constructor(public params: NavParams, public viewCtrl: ViewController, public mapSvc: MapService) {
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
    this.mapSvc.addPin(this.lat, this.lng, this.title, this.notes, this.mapId);
    this.viewCtrl.dismiss();
  }
}
