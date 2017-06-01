(function (w, objname) {
	var Conversion = function () {}
	w[objname] = new Conversion();

	var x_pi = 3.14159265358979324 * 3000.0 / 180.0;
	var pi = 3.14159265358979324;
    var a = 6378245.0;
    var ee = 0.00669342162296594323;

	var LatLng = function (lat, lng) {
		this.lat = lat ? parseFloat(lat) : 0.0;
		this.lng = lng ? parseFloat(lng) : 0.0;
	}
	LatLng.prototype.toFixed = function (num) {
		this.lat = parseFloat(this.lat.toFixed(num));
		this.lng = parseFloat(this.lng.toFixed(num));
		return this;
	}
	Conversion.prototype.latlng = function (lat, lng) {
		return new LatLng(lat, lng);
	}

    Conversion.prototype.wgs_gcj = function (wgLoc) {
        if (outOfChina(wgLoc.lat, wgLoc.lng)) return wgLoc;
        var dLat = transformLat(wgLoc.lng - 105.0, wgLoc.lat - 35.0);
        var dLon = transformLon(wgLoc.lng - 105.0, wgLoc.lat - 35.0);
        var radLat = wgLoc.lat / 180.0 * pi;
        var magic = Math.sin(radLat);
        magic = 1 - ee * magic * magic;
        var sqrtMagic = Math.sqrt(magic);
        dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * pi);
        dLon = (dLon * 180.0) / (a / sqrtMagic * Math.cos(radLat) * pi);
        return new LatLng(wgLoc.lat + dLat, wgLoc.lng + dLon);
    }

    Conversion.prototype.gcj_wgs = function (gcLoc) {
        var wgLoc = new LatLng(gcLoc.lat, gcLoc.lng);
        var dLoc = new LatLng();
        while (true) {
            var currGcLoc = this.wgs_gcj(wgLoc);
            dLoc.lat = gcLoc.lat - currGcLoc.lat;
            dLoc.lng = gcLoc.lng - currGcLoc.lng;
            if (Math.abs(dLoc.lat) < 1e-7 && Math.abs(dLoc.lng) < 1e-7) {
                return wgLoc;
            }
            wgLoc.lat += dLoc.lat;
            wgLoc.lng += dLoc.lng;
        }
    }

    Conversion.prototype.gcj_bd = function (gcLoc) {
        var x = gcLoc.lng, y = gcLoc.lat;
        var z = Math.sqrt(x * x + y * y) + 0.00002 * Math.sin(y * x_pi);
        var theta = Math.atan2(y, x) + 0.000003 * Math.cos(x * x_pi);
        return new LatLng(z * Math.sin(theta) + 0.006, z * Math.cos(theta) + 0.0065);
    }

    Conversion.prototype.bd_gcj = function (bdLoc) {
        var x = bdLoc.lng - 0.0065, y = bdLoc.lat - 0.006;
        var z = Math.sqrt(x * x + y * y) - 0.00002 * Math.sin(y * x_pi);
        var theta = Math.atan2(y, x) - 0.000003 * Math.cos(x * x_pi);
        return new LatLng(z * Math.sin(theta), z * Math.cos(theta));
    }

    Conversion.prototype.wgs_bd = function (ll) {
        return this.gcj_bd(this.wgs_gcj(ll));
    }

    Conversion.prototype.bd_wgs = function (ll) {
        return this.gcj_wgs(this.bd_gcj(ll));
    }

    var outOfChina = function (lat, lon) {
        if (lon < 72.004 || lon > 137.8347) return true;
        if (lat < 0.8293 || lat > 55.8271) return true;
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
})(global, 'convert')

const lat = '28.89', lng = 120.68;
const print = console.log;
print('bd_gcj', convert.bd_gcj(convert.latlng(lat, lng)).toFixed(6));
print('bd_wgs', convert.bd_wgs(convert.latlng(lat, lng)).toFixed(6));
print('gcj_bd', convert.gcj_bd(convert.latlng(lat, lng)).toFixed(6));
print('gcj_wgs', convert.gcj_wgs(convert.latlng(lat, lng)).toFixed(6));
print('wgs_bd', convert.wgs_bd(convert.latlng(lat, lng)).toFixed(6));
print('wgs_gcj', convert.wgs_gcj(convert.latlng(lat, lng)).toFixed(6));
