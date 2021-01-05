# Facebook login

`
POST /login/device-based/regular/login/ HTTP/1.1
Host: mbasic.facebook.com
User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
Accept: */*;
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate
Content-Type: application/x-www-form-urlencoded
Content-Length: 187
Origin: https://mbasic.facebook.com
Connection: close
Cookie: sb=DoPxX9e3gacWieXay1gIPwzM; datr=DoPxX-wEhsyAtYDGcSnMs3cG;
Upgrade-Insecure-Requests: 1

lsd=AVpJfDk-quE&jazoest=2958&m_ts=1609786134&li=FmPzX868_5Lji6O7tYWN6B8-&try_number=0&unrecognized_tries=0&email=ghzant.y%40gmail.com&pass=password%28password%29&login=Iniciar+sesi%C3%B3n
`

* Enter email: ghzant.y@gmail.com -> ghzant.y@40gmail.com
* Enter password: password(password) -> password%28password%29password
* sb
* datr

### cookie/consent/
* fr=123LZQ8aVPeFJDeci.AWWS0hTNq7xcYuNyPYBifuuU1aY.Bf8YMO.6P.AAA.0.0.Bf8Z25.AWXDJ1eMQgQ;
* sb=DoPxX9e3gacWieXay1gIPwzM;
* datr=DoPxX-wEhsyAtYDGcSnMs3cG;
* wd=950x943;
* locale=es_LA

### Set-Cookie
* sb=DoPxX9e3gacWieXay1gIPwzM; expires=Tue, 03-Jan-2023 11:23:05 GMT; Max-Age=63072000; 
* c_user=100044781015699; expires=Mon, 03-Jan-2022 11:23:04 GMT; Max-Age=31535999; 
* xs=10%3Aci_DHAlSfo2ryw%3A2%3A1609672985%3A-1%3A-1; expires=Mon, 03-Jan-2022 11:23:04 GMT; Max-Age=31535999; 
* fr=123LZQ8aVPeFJDeci.AWVuAqrOly7W4jn1j0Ukt5M_wyY.Bf8YMO.6P.AAA.0.0.Bf8akZ.AWW6zVs0IlU; expires=Sat, 03-Apr-2021 11:23:03 GMT; Max-Age=7775998; 

### POST /ajax/qm/...
Set-Cookie:
* spin=r.1003142924_b.trunk_t.1609673721_s.1_v.2_; expires=Mon, 04-Jan-2021 12:35:21 GMT; Max-Age=90000;

- added at each sentence: path=/; domain=.facebook.com; secure; httponly; SameSite=None

# Facebook Reactions
`
POST /api/graphql/ HTTP/1.1
Host: www.facebook.com
User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
Accept: */*
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate
Content-Type: application/x-www-form-urlencoded
Content-Length: 1388
Origin: https://www.facebook.com
Connection: close
Referer: https://www.facebook.com/IMSSmx/photos/a.10150721909064578/10160673567979578/
Cookie: fr=123LZQ8aVPeFJDeci.AWWQKqsU4PloCriWUAlqHYnw6Ck.Bf8YMO.6P.AAA.0.0.Bf8lKi.AWU7Xe_eT-U; sb=DoPxX9e3gacWieXay1gIPwzM; datr=DoPxX-wEhsyAtYDGcSnMs3cG; wd=1920x947; locale=es_LA; c_user=100044781015699; xs=10%3A10zzeAvKHseC8w%3A2%3A1609712759%3A-1%3A-1%3A%3AAcV5ExHCl0eFjCvubhQU33u6sNoA8xo4IYnq5Apyhg; spin=r.1003143044_b.trunk_t.1609712761_s.1_v.2_; presence=C%7B%22t3%22%3A%5B%5D%2C%22utc3%22%3A1609714417766%2C%22v%22%3A1%7D

av=100044781015699&__user=100044781015699&__a=1&__dyn=7AzHxqU5a5Q2m3mbG2KnFw9uu2i5U4e0yoW3q322aewXwnEbotwp8O2S1DwUx609vCxS320om78-0BE88427Uy11xmfz83WwgEcHzoaEaoG0Boy1PwBgK7qwpE31wnEcUC68gwHwxwQzXxG1Pxi4UaEW0D8qBwJK5Umxm5oe8aUlxfxmu3W2i4U72dG5Ey2a2-&__csr=g8a3svaDsuxlsZgxt49FZGTtnW99KrqDmG-XHSSaXsCi8S9tuAJaFLZpkV9R8i9G8VrK_KhkmcDxuGubGGB-injWhaAzAIxoCdDGeBwyDx91e2Gbx2UyWy-ex-5VEqUS6qDyFrBxiayorx6heiqEC6898kF1a3Ciby98uxOqiErCg8oybCwNzUmz888oxu1xJ3ooye26iEowGxe11BUmwDwWx-m3yi48O6E98kxS3-2WdwUwOwRw0TIxkw0lxw0GJw288720gC08dK09gwr6ErguwuEgjws9Udd0u81do7Etu1bwxwEUbo3Iw2uo0kmzo&__req=1f&__beoa=0&__pc=EXP2%3Acomet_pkg&dpr=1&__ccg=GOOD&__rev=1003143062&__s=e52cd3%3Atl3wgc%3Aw1h268&__hsi=6913676297949371087-0&__comet_req=1&__comet_env=fb&fb_dtsg=AQFHH9FxCBjq%3AAQEz72ZLRg6v&jazoest=21982&__spin_r=1003143062&__spin_b=trunk&__spin_t=1609715702&fb_api_caller_class=RelayModern&fb_api_req_friendly_name=CometUFIFeedbackReactMutation&variables=%7B%22input%22%3A%7B%22feedback_id%22%3A%22ZmVlZGJhY2s6MTAxNjA2NzM1NjgxMDQ1Nzg%3D%22%2C%22feedback_reaction%22%3A4%2C%22feedback_source%22%3A%22MEDIA_VIEWER%22%2C%22is_tracking_encrypted%22%3Atrue%2C%22tracking%22%3A%5B%5D%2C%22session_id%22%3A%2252a455f8-c745-448a-abe0-669dab6d6657%22%2C%22actor_id%22%3A%22100044781015699%22%2C%22client_mutation_id%22%3A%224%22%7D%2C%22useDefaultActor%22%3Afalse%7D&server_timestamps=true&doc_id=4136721883019636
`

* variables= ( URL encoded ) -> (decode) -> (JSON) -> 
variables={"input":{"feedback_id":"ZmVlZGJhY2s6MTAxNjA2NzM1NjgxMDQ1Nzg=","feedback_reaction":4,"feedback_source":"MEDIA_VIEWER","is_tracking_encrypted":true,"tracking":[],"session_id":"52a455f8-c745-448a-abe0-669dab6d6657","actor_id":"100044781015699","client_mutation_id":"4"},"useDefaultActor":false}

- variables.input.feedback_reaction -> {
	  1:like,
	  2:heart,
	  3:me-importa,
	  4:funny,
	  5:asombra,
	  6:sad,
	  7:angry
}

- variables.input.actor_id = document.cookie.c_user = 100044781015699 (check Set-Cookie ^^^) 

