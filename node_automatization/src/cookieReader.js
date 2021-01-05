import requestHeaders from "./requestHeaders.js";
import tough_cookie_pkg from 'tough-cookie';
const {Cookie} = tough_cookie_pkg;

const cookieReader = {
	read: (responseHeaders) => {
		let cookies;
		if (responseHeaders['set-cookie'] instanceof Array)
			cookies = responseHeaders['set-cookie'].map(Cookie.parse);
		else
			cookies = [Cookie.parse(responseHeaders['set-cookie'])];
		return cookies;
	},
	update: (uid,cookies) => {
		cookieReader.records[uid] = cookies;
	},
	filterArray:["datr","sb","c_user","xs","fr"],
	records:{}
}

export default cookieReader;
