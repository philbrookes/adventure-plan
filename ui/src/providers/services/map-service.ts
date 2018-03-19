import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';


export interface MapPointMetadata {
  title: string;
  notes: string;
}

export interface MapMetadata {
  title: string;
}

export interface MapPoint {
  latitude: number;
  longitude: number;
  metadata: MapPointMetadata;
}
export interface Map {
  id: number;
  center: MapPoint;
  zoom: number;
  metadata: MapMetadata;
  interests: MapPoint[];
}
export interface MapsResponse {
  maps: Map[];
}
@Injectable()
export class MapService {

  constructor(public http: HttpClient) {
    
  }

  getMaps(): Promise<Map[]> {
    return new Promise((resolve, reject) => {
      this.http.get('http://localhost:8080/maps').subscribe((mapData: MapsResponse) => {
        resolve(mapData.maps as Map[] || [])
      });
    });
  }

  addPin(lat, lng, title, notes, mapId): Promise<MapPoint> {
    let options = {headers: {'Content-Type': 'application/x-www-form-urlencoded'}};
    //let body = "latitude=" + lat +"&longitude=" + lng + "&title=" + title + "&notes=" + notes;
    let body = {latitude: lat, longitude: lng, metadata: {title: title, notes: notes}};
    
    return new Promise((resolve, reject) => {
      this.http.post('http://localhost:8080/maps/' + mapId + "/pin", body, options).subscribe((point: MapPoint) => {
        resolve(point as MapPoint);
      });
    });
  }
}
