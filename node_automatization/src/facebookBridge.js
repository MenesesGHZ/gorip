import got from 'got';
import cookieReader from "./cookieReader.js";
import requestHeaders from  './requestHeaders.js';

const facebookBridge = {
	 login: async(email,password,uid) => {
		const headers = cookieReader._cookieIntoHeader(requestHeaders,cookieUID[uid]);
		const parameters = [[]];
		const instance = got.extend({
			"prefixUrl":facebookBridge.prefixURL,
			"headers":requestHeaders
		});
		const response = await instance.post("/login/device-based/regular/login/").json();
		console.log(response);

	},
	makeReaction:(uid,reaction,url) => {
	
	},
	get: async(uid,suffix) => {
		const response = await got(facebookBridge.prefixURL+suffix);
		const cookie = cookieReader.read(response.headers);
		console.log(cookie)
		cookieReader.update(uid,cookie);
		console.log(cookieReader.records);
	},
	prefixURL:"https://mbasic.facebook.com"  
}

facebookBridge.get("100044781015699","/");
