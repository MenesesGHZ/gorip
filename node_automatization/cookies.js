const {promisify} = require('util');
const got = require('got');
const {CookieJar} = require('tough-cookie');

const headers = {
	GET:{
	"Host": "mbasic.facebook.com",
	"User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0",
	"Accept": "*/*",
	"Accept-Language": "en-US,en;q=0.5",
	"Accept-Encoding": "gzip, deflate",
	"Connection": "close",
	"Content-Length":"2"
	},
	POST:{
		/**/
	}
}


let cookieUID = {};
const facebookPrefixURL ='https://mbasic.facebook.com';  

const cookieReader = {
	facebookPage: async(uid) => {
		const response = await got(facebookPrefixURL);
		const setCookie = cookieReader._getCookieHeader(response.headers);
		cookieUID[uid] = {"datr":setCookie["datr"],"sb":setCookie["sb"]};
	},
	facebookLoginPage: async(email,password,uid){
		const request_headers = cookieReader._cookieIntoHeader(headers,cookieUID[uid]);
		const instance = got.extend({
			"prefixUrl":facebookPrefixURL,
			"headers":request_headers
		});
		const response_headers = await instance.post().json();
		console.log(response_headers);
	},

	_cookieIntoHeader:(headers,cookie){
		let cookieString = "";
		for(let keys of Object.keys(cookie)){
			cookieString += `${key}=`+`${cookie[key]};`
		}
		headers["Cookie"] = cookieString;
		return headers;
	},
	_getCookieHeader: (headers) => {
		let setCookiePairs = "".concat(...headers["set-cookie"].map(val=>val+";")),
			setCookie={};
		setCookiePairs = setCookiePairs.split(";").map(val=>val.trim().split("=")).slice(0,-1)
		for(let pair of setCookiePairs)
			setCookie[pair[0]] = pair[1];
		return setCookie;
	}
}


cookieReader.facebookPage("mock-user");
