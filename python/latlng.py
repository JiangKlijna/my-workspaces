import math as Math

Math.abs = lambda x: x if x >= 0 else -x

pi = 3.14159265358979324
a = 6378245.0
ee = 0.00669342162296594323
x_pi = 3.14159265358979324 * 3000.0 / 180.0


class LatLon:
    def __init__(self, lat=0.0, lng=0.0):
        self.lat = lat
        self.lng = lng

    def __str__(self):
        return "LatLon{lat = %s, lon = %s}" % (self.lat, self.lng)


class Convert:
    def wgs_gcj(self, latlon): pass

    def wgs_bd(self, latlon): pass

    def gcj_wgs(self, latlon): pass

    def gcj_bd(self, latlon): pass

    def bd_gcj(self, latlon): pass

    def bd_wgs(self, latlon): pass


class ConvertImpl(Convert):

    def __init__(self):
        self.bd_wgs = lambda latlon: self.gcj_wgs(self.bd_gcj(latlon))
        self.wgs_bd = lambda latlon: self.gcj_bd(self.wgs_gcj(latlon))

    def wgs_gcj(self, latlon):
        if self.outOfChina(latlon.lat, latlon.lng): return latlon
        dLat = self.transformLat(latlon.lng - 105.0, latlon.lat - 35.0)
        dLon = self.transformLon(latlon.lng - 105.0, latlon.lat - 35.0)
        radLat = latlon.lat / 180.0 * pi
        magic = Math.sin(radLat)
        magic = 1 - ee * magic * magic
        sqrtMagic = Math.sqrt(magic)
        dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * pi)
        dLon = (dLon * 180.0) / (a / sqrtMagic * Math.cos(radLat) * pi)
        return LatLon(latlon.lat + dLat, latlon.lng + dLon)

    def gcj_wgs(self, latlon):
        wgLoc = LatLon(latlon.lat, latlon.lng)
        dLoc = LatLon()
        while True:
            currGcLoc = self.wgs_gcj(wgLoc)
            dLoc.lat = latlon.lat - currGcLoc.lat
            dLoc.lng = latlon.lng - currGcLoc.lng
            if Math.abs(dLoc.lat) < 1e-7 and Math.abs(dLoc.lng) < 1e-7:
                return wgLoc
            wgLoc.lat += dLoc.lat
            wgLoc.lng += dLoc.lng

    def gcj_bd(self, latlon):
        x = latlon.lng
        y = latlon.lat
        z = Math.sqrt(x * x + y * y) + 0.00002 * Math.sin(y * x_pi)
        theta = Math.atan2(y, x) + 0.000003 * Math.cos(x * x_pi)
        return LatLon(z * Math.sin(theta) + 0.006, z * Math.cos(theta) + 0.0065)

    def bd_gcj(self, latlon):
        x = latlon.lng - 0.0065
        y = latlon.lat - 0.006
        z = Math.sqrt(x * x + y * y) - 0.00002 * Math.sin(y * x_pi)
        theta = Math.atan2(y, x) - 0.000003 * Math.cos(x * x_pi)
        return LatLon(z * Math.sin(theta), z * Math.cos(theta))

    def outOfChina(self, lat, lon):
        if lon < 72.004 or lon > 137.8347: return True
        if lat < 0.8293 or lat > 55.8271: return True
        return False

    def transformLat(self, x, y):
        ret = -100.0 + 2.0 * x + 3.0 * y + 0.2 * y * y + 0.1 * x * y + 0.2 * Math.sqrt(x if x > 0 else -x)
        ret += (20.0 * Math.sin(6.0 * x * pi) + 20.0 * Math.sin(2.0 * x * pi)) * 2.0 / 3.0
        ret += (20.0 * Math.sin(y * pi) + 40.0 * Math.sin(y / 3.0 * pi)) * 2.0 / 3.0
        ret += (160.0 * Math.sin(y / 12.0 * pi) + 320 * Math.sin(y * pi / 30.0)) * 2.0 / 3.0
        return ret

    def transformLon(self, x, y):
        ret = 300.0 + x + 2.0 * y + 0.1 * x * x + 0.1 * x * y + 0.1 * Math.sqrt(x if x > 0 else -x)
        ret += (20.0 * Math.sin(6.0 * x * pi) + 20.0 * Math.sin(2.0 * x * pi)) * 2.0 / 3.0
        ret += (20.0 * Math.sin(x * pi) + 40.0 * Math.sin(x / 3.0 * pi)) * 2.0 / 3.0
        ret += (150.0 * Math.sin(x / 12.0 * pi) + 300.0 * Math.sin(x / 30.0 * pi)) * 2.0 / 3.0
        return ret


Convert = ConvertImpl()

if __name__ == '__main__':
    lat, lng = 28.89, 120.68
    print('bd_gcj', Convert.bd_gcj(LatLon(lat, lng)))
    print('bd_wgs', Convert.bd_wgs(LatLon(lat, lng)))
    print('gcj_bd', Convert.gcj_bd(LatLon(lat, lng)))
    print('gcj_wgs', Convert.gcj_wgs(LatLon(lat, lng)))
    print('wgs_bd', Convert.wgs_bd(LatLon(lat, lng)))
    print('wgs_gcj', Convert.wgs_gcj(LatLon(lat, lng)))
