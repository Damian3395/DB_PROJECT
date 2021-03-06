/* JKM, cookies, cookies-jkm.js, 1-5-06 */

function reldate(days) {
var d;
d = new Date();

/* We need to add a relative amount of time to
the current date. The basic unit of JavaScript
time is milliseconds, so we need to convert the
days value to ms. Thus we have
ms/day
= 1000 ms/sec *  60 sec/min * 60 min/hr * 24 hrs/day
= 86,400,000. */

d.setTime(d.getTime() + days*86400000);
return d.toGMTString();
}

function readCookie(name) {
var s = document.cookie, i;
if (s)
for (i=0, s=s.split('; '); i<s.length; i++) {
s[i] = s[i].split('=', 2);
if (unescape(s[i][0]) == name)
return unescape(s[i][1]);
}
return null;
}

function makeCookie(name, value, p) {
var s, k;
s = escape(name) + '=' + escape(value);
if (p) for (k in p) {

/* convert a numeric expires value to a relative date */

if (k == 'expires')
p[k] = isNaN(p[k]) ? p[k] : reldate(p[k]);

/* The secure property is the only special case
here, and it causes two problems. Rather than
being '; protocol=secure' like all other
properties, the secure property is set by
appending '; secure', so we have to use a
ternary statement to format the string.

The second problem is that secure doesn't have
any value associated with it, so whatever value
people use doesn't matter. However, we don't
want to surprise people who set { secure: false }.
For this reason, we actually do have to check
the value of the secure property so that someone
won't end up with a secure cookie when
they didn't want one. */

if (p[k])
s += '; ' + (k != 'secure' ? k + '=' + p[k] : k);
}
document.cookie = s;
return readCookie(name) == value;
}

function rmCookie(name) {
return !makeCookie(name, '', { expires: -1 });
}
