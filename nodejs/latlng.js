(function (w, objname) {
	var Conversion = function () {}
	var c = new Conversion();
	
	c.LatLng = function (lat, lng) {
		if (lat && lng) {
			this.lat = lat;
			this.lng = lng;
		} else {
			this.lat = 0;
			this.lng = 0;
		}
	}
	
	w[objname] = c;
	
	var x_pi = 3.14159265358979324 * 3000.0 / 180.0;
	var pi = 3.14159265358979324;
    var a = 6378245.0;
    var ee = 0.00669342162296594323;
	
	//百度->gps
    Conversion.prototype.bd_wgs = function (latlng) {
        return this.gcj_wgs(this.bd_gcj(latlng));
    }
	
	//百度->火星
    Conversion.prototype.bd_gcj = function (latlng) {
        var x = latlng.lng - 0.0065, y = latlng.lat - 0.006;
        var z = Math.sqrt(x * x + y * y) - 0.00002 * Math.sin(y * x_pi);
        var theta = Math.atan2(y, x) - 0.000003 * Math.cos(x * x_pi);
        return new c.LatLng(z * Math.sin(theta), z * Math.cos(theta));
    }
	
	//火星->gps
    Conversion.prototype.gcj_wgs = function (gcLoc) {
        var wgLoc = new c.LatLng();
        wgLoc.lat = gcLoc.lat;
        wgLoc.lng = gcLoc.lng;

        var currGcLoc = new c.LatLng();
        var dLoc = new c.LatLng();
        while (true) {
            currGcLoc = this.wgs_gcj(wgLoc);
            dLoc.lat = gcLoc.lat - currGcLoc.lat;
            dLoc.lng = gcLoc.lng - currGcLoc.lng;
            if (Math.abs(dLoc.lat) < 1e-7 && Math.abs(dLoc.lng) < 1e-7) {
                return wgLoc;
            }
            wgLoc.lat += dLoc.lat;
            wgLoc.lng += dLoc.lng;
        }
    }

	//火星->百度
    Conversion.prototype.gcj_bd = function (gcLoc) {
        var x = gcLoc.lng, y = gcLoc.lat;
        var z = Math.sqrt(x * x + y * y) + 0.00002 * Math.sin(y * x_pi);
        var theta = Math.atan2(y, x) + 0.000003 * Math.cos(x * x_pi);
        return new c.LatLng(z * Math.sin(theta) + 0.006, z * Math.cos(theta) + 0.0065);
    }

    //gps->百度
    Conversion.prototype.wgs_bd = function (latlng) {
        return this.gcj_bd(this.wgs_gcj(latlng));
    }
	
	//gps->火星
    Conversion.prototype.wgs_gcj = function (latlng) {
		var lat = latlng.lat, lng = latlng.lng;
        if (outOfChina(lat, lng)) {
            return new c.LatLng(lat, lng);
        }
        var dLat = transformLat(lng - 105.0, lat - 35.0);
        var dLon = transformLon(lng - 105.0, lat - 35.0);
        var radLat = lat / 180.0 * pi;
        var magic = Math.sin(radLat);
        magic = 1 - ee * magic * magic;
        var sqrtMagic = Math.sqrt(magic);
        dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * pi);
        dLon = (dLon * 180.0) / (a / sqrtMagic * Math.cos(radLat) * pi);
        return new c.LatLng(lat + dLat, lng + dLon);
    }

    var outOfChina = function (lat, lng) {
        if (lat < 72.004 || lat > 137.8347)
            return true;
        if (lng < 0.8293 || lng > 55.8271)
            return true;
        return false;
    }

    var transformLat = function (x, y) {
        var ret = -100.0 + 2.0 * x + 3.0 * y + 0.2 * y * y + 0.1 * x * y + 0.2 * Math.sqrt(x > 0 ? x : -x);
        ret += (20.0 * Math.sin(6.0 * x * pi) + 20.0 * Math.sin(2.0 * x * pi)) * 2.0 / 3.0;
        ret += (20.0 * Math.sin(y * pi) + 40.0 * Math.sin(y / 3.0 * pi)) * 2.0 / 3.0;
        ret += (160.0 * Math.sin(y / 12.0 * pi) + 320 * Math.sin(y * pi / 30.0)) * 2.0 / 3.0;
        return ret;
    }

    var transformLon = function (x, y) {
        var ret = 300.0 + x + 2.0 * y + 0.1 * x * x + 0.1 * x * y + 0.1 * Math.sqrt(x > 0 ? x : -x);
        ret += (20.0 * Math.sin(6.0 * x * pi) + 20.0 * Math.sin(2.0 * x * pi)) * 2.0 / 3.0;
        ret += (20.0 * Math.sin(x * pi) + 40.0 * Math.sin(x / 3.0 * pi)) * 2.0 / 3.0;
        ret += (150.0 * Math.sin(x / 12.0 * pi) + 300.0 * Math.sin(x / 30.0 * pi)) * 2.0 / 3.0;
        return ret;
    }
})(global, "conversion")

const lat = 120.68, lng = 28.89;
const print = console.log;
print("bd_gcj", conversion.bd_gcj(new conversion.LatLng(lat, lng)));
print("bd_wgs", conversion.bd_wgs(new conversion.LatLng(lat, lng)));
print("gcj_bd", conversion.gcj_bd(new conversion.LatLng(lat, lng)));
print("gcj_wgs", conversion.gcj_wgs(new conversion.LatLng(lat, lng)));
print("wgs_bd", conversion.wgs_bd(new conversion.LatLng(lat, lng)));
print("wgs_gcj", conversion.wgs_gcj(new conversion.LatLng(lat, lng)));