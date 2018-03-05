import { Component } from '@angular/core';
import { NavController } from 'ionic-angular';
import { MapService, Map } from '../../providers/services/map-service'
import { HomePage } from '../../pages/home/home';

@Component({
  selector: 'page-list',
  templateUrl: 'list.html'
})
export class ListPage {
  maps: Map[];
  nav: NavController;

  constructor(public navCtrl: NavController, public mapSvc: MapService) {
    this.nav = navCtrl;
    mapSvc.getMaps().then(maps => {
      this.maps = maps;
    })
    .catch(error => {

    });
  }

  mapClicked(event, map){
    this.nav.push(HomePage, {map: map});
  }
}
