import { Component, ViewChild, ElementRef } from '@angular/core';
import { ModalController, NavController, NavParams, Events } from 'ionic-angular';
import leaflet from 'leaflet';
import { Map } from '../../providers/services/map-service';
import { PinEdit } from '../pin-edit/pin-edit';

@Component({
  selector: 'page-home',
  templateUrl: 'home.html'
})
export class HomePage {
  loadedMap: boolean = false;
  @ViewChild('map') mapContainer: ElementRef;
  map: any;
  mapData: Map;
  modals: ModalController;
  interests: leaflet.featureGroup;

  constructor(public navCtrl: NavController, public navParams: NavParams, public modalCtrl: ModalController, public events: Events) {
    this.mapData = <Map>navParams.get("map");
    this.modals = modalCtrl;
    this.interests = leaflet.featureGroup();
    this.events.subscribe("interests:new", (point) => {
      let intMarker: any = new leaflet.marker([point.latitude, point.longitude]);
      intMarker.bindPopup("<div class='marker-popup'><h3>" + point.metadata.title + "</h3><p>" + point.metadata.notes + "</p></div>");
      this.interests.addLayer(intMarker);
    });
    this.events.subscribe("map:updated", (map) => {
      if(map.id === this.mapData.id){
        this.mapData = map;
        this.loadmap();
      }
    })
  }

  ionViewDidLoad() {
    this.loadmap();
  }
 
  mapClick(event) {
    let pinModal = this.modals.create(PinEdit, { lat: event.latlng.lat, lng: event.latlng.lng, mapid: this.mapData.id });
    pinModal.present();
  }


  loadmap() {
    this.map = leaflet.map("map").fitWorld();
    this.map.on('click', this.mapClick.bind(this));
    leaflet.tileLayer('http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {maxZoom: 18}).addTo(this.map);
    if(this.mapData){
      this.map.setView(new leaflet.LatLng(this.mapData.center.latitude, this.mapData.center.longitude), this.mapData.zoom);
      this.drawInterests();
    }
  }

  drawInterests() {
    this.map.removeLayer(this.interests);
    this.mapData.interests.forEach(interest => {
      let intMarker: any = new leaflet.marker([interest.latitude, interest.longitude]);
      intMarker.bindPopup("<div class='marker-popup'><h3>" + interest.metadata.title + "</h3><p>" + interest.metadata.notes + "</p></div>");
      this.interests.addLayer(intMarker);
    });
    this.map.addLayer(this.interests);
  }

}
