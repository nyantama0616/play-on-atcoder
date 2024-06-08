#include <bits/stdc++.h>

#pragma once

using namespace std;

using ll = int64_t;
using ull = unsigned long long;
using ld = long double;
using P = pair<ll,ll>;
using V = vector<ll>;
using M = vector<vector<ll>>;

#define rep(i, n) for (ll i = 0; i < static_cast<ll>(n); ++i)
#define rep_rev(i, n) for (ll i = n - 1; i >= 0; --i)
#define rep2(i, a, b) for (ll i = a; i < static_cast<ll>(b); ++i)
#define rep_rev2(i, a, b) for (ll i = b - 1; i >= a; --i)
#define all(v) (v).begin(), (v).end()
#define bit(n,k) (((n)>>(k))&1)
#define isin(x, l, r) (l <= (x) && (x) < r)
#define dx4 {-1, 0, 1, 0}
#define dy4 {0, 1, 0, -1}
#define dx8 {-1, -1, 0, 1, 1, 1, 0, -1}
#define dy8 {0, 1, 1, 1, 0, -1, -1, -1}
#define length(a) static_cast<ll>(a.size())
#define get(x, i) get<i>(x)
#define sign(x) ((x == 0) ? 0 : (x) / abs(x))
#define ndigits(x) static_cast<ll>(floor(log10(x)) + 1)
#define deg2rad(x) ((x) * M_PI / 180)
#define rad2deg(x) ((x) * 180 / M_PI)
#define log(x, y) (log(y) / log(x))
#define scanvar(T, x) T x; cin >> x
#define scanll(x) ll x; cin >> x
#define scanv(T, v, n) vector<T> v; for (int i = 0; i < n; ++i) { T temp; cin >> temp; (v).push_back(temp); }
#define scanvs(v) vector<char> v; do { scanvar(string, temp); for (char c : temp) {v.push_back(c); } } while (false)
#define scanm(T, m, a, b) vector<vector<T>> m(a, vector<T>(b)); rep(i, a) { rep(j, b) { cin >> m[i][j]; } }
#define scanms(m, a, b) vector<vector<char>> m(a, vector<char>(b)); rep(i,a) { scanvar(string, temp); rep(j, b) { m[i][j] = temp[j]; } }
#define pl(x) cout << (x) << endl
#define printld(x) printf("%.12Lf\n", x)
#define printv(v) for (auto __x: v) {cout << __x << " ";} cout << endl;
#define printm(m) for (auto __v : m) { printv(__v); }

ll _pow(ll a, ll b) { ll res = 1; while (b > 0) { if (b & 1) res *= a; a *= a; b >>= 1; } return res; }
template <class T1, class T2, class Func> void __map(vector<T1> v, vector<T2>& res, Func f) { int n = v.size(); rep(i, n) { res[i] = f(v[i]); } }
template <class T> void __unique(vector<T>& v) { auto p = unique(all(v)); v.erase(p, v.end()); }
template <class T> void __join(const vector<T> v, const string separator, string& res) { ll n = v.size(); rep(i, n - 1) { res += to_string(v[i]) + separator; } res += to_string(v[n - 1]); }
void __joins(const vector<char> v, const string separator, string& res) { ll n = v.size(); rep(i, n - 1) { res += v[i] + separator; } if (n > 0) res += v[n - 1]; }
void __to_string(ll x, ll num, string& res, ll pad = -1) { while (x >= num) { res = to_string(x % num) + res; x /= num; } if (x > 0) res = to_string(x) + res; ll d = pad - res.size(); rep(i, d) { res = '0' + res; } }

// Debug
#define DEBUG_MODE 1
#define debug(x) if (DEBUG_MODE) { cout << "  \033[35m" << #x << "\033[m" << ": " << x << endl; }
#define debugld(x) if (DEBUG_MODE) { cout << `"  \033[35m`" << #x << `"\033[m`" << `": `"; printld(x);}
#define debugv(v) if (DEBUG_MODE) { cout << "  \033[35m" << #v << "\033[m" << ": "; for (auto& __x : v) {cout << __x << " "; } cout << endl; }
#define debugm(m) if (DEBUG_MODE) { cout << "  \033[35m" << #m << "\033[m" << ":\n"; for (auto __v : m) { cout << "  "; printv(__v); } cout << endl; }
#if DEBUG_MODE

#endif
