const {promisify} = require('util');
const got = require('got');
const {CookieJar} = require('tough-cookie');

const headers = {
	"Host": "mbasic.facebook.com",
	"User-Agent": "Mozilla/5.0",
	"Accept": "*/*",
	"Accept-Language": "en-US,en;q=0.5",
	"Accept-Encoding": "gzip, deflate",
	"Connection": "close",
	"Content-Length":"2"
}
let cookie_uid = {};
const facebookPrefixURL ='https://mbasic.facebook.com/';  

const cookieReader = {
	facebookPage: async(uid) => {
		const response = await got(facebookPrefixURL);
		const setCookie = cookieReader._getCookieHeader(response.headers);
		console.log(setCookie);
	},
	
	_getCookieHeader: (headers) => {
		let setCookie = "".concat(...headers["set-cookie"].map(val=>val+";"))
		setCookie = setCookie.split(";").map(val=>val.trim().split("=")).slice(0,-1)
		return setCookie;
	}

}


cookieReader.facebookPage()
