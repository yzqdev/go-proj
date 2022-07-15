package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	var url = `https://hk4e-api.mihoyo.com/event/gacha_info/api/getGachaLog?authkey_ver=1&sign_type=2&auth_appid=webview_gacha&init_type=301&gacha_id=e432d80a79eb2a8953331fe921c7b6dcf25aa808&timestamp=1657669515&lang=zh-cn&device_type=pc&ext=%7b%22loc%22%3a%7b%22x%22%3a2942.8603515625%2c%22y%22%3a274.06182861328127%2c%22z%22%3a-1247.0826416015625%7d%2c%22platform%22%3a%22WinST%22%7d&game_version=CNRELWin2.8.0_R8589120_S8812249_D8644052&plat_type=pc&region=cn_gf01&authkey=ysM7%2fOM1%2fiJwMXUCc7s8mpLbFD4MENVnwRVFo369EOvbHK%2fy92KC%2fUnycV%2fIx26c%2fVF21xO4BouA6M6qCvDp9AMtt6GQj6DV28zXBsrZHxh4PocAT564uynHfXz4PvQ6lfQizgtVITNOt6mUCCS4WKZfKTjcpCDYFMXARIy01CeagNwoutdsNQwIc9Qpi87KjAwsUpFUfXB6pWCwHyeoqf29AuO4w0%2fNFKRbVDGFOBtBQ%2fIz4n4jJLkQPHzpjDNOS5tqOf11zBnAKjGaQxu1VrXWlK3kF1NGdJgmHy%2bIuHwxRZxpvlsC%2b3MxVm0vz5Sytfl5JpHX1jhn9bFMIR5Ytns0AobwP%2bWsADCEOcZiUelYXWgYjs19E0e7GH8KhkR8J4xvfkuVt0V9Exh4Iho%2fdsJp4K8Rgxv0DJmnNhYawQAmoKiv5QxrsRp%2f9C1s04rFaVw1iaY2lq85GtjbG59CWU4QB5puJaalqhEZwzYRENiPujgQpyxhgVmEjy%2fMGW7vuwAuzDgn01fCDcJ%2fJhrSIy7rJNEMf6W313z8LS54kTIsGKYDNePHU8%2bkGHB4lRCofvjj66uBEw5aarI%2b1la%2b1OI%2fZaBgXBZ4kV%2bnA67m9%2fQ95M5ADXqkCHlN8c1EA2P5%2fDG8LY6OtBzENBnsUYtWs9bvT8UDTISu78%2bPfXhNN7m9pXoSIO0aJa0LDoOk%2bIKGYoFV3r0DuTeVfEpRHWmeCq%2btSPO6IL%2fg%2fNhK76Vn2V3%2bvaYURGXJyG1nVIshx7TR8P05m3z4nzTfwyXmJy%2bYda8sYdFRRnop8yP9oMm9I1GPqmJmV%2bE9qzHcgLxfrRB5idrBzbeOWnHUQ0XA43cHW1nz8%2f%2bfEoAt3rYjSF7mW%2fWjU6bp0cpkyk0AsUgOMsBzfQ%2fckU4SPj3ytL4az9rpO60N3dv%2bk0kw6uwIHuUk%2flLkSyZDpAS8O1arXFAWpEr8FczPvDQm9wLBlBFbwcQ2ZaRb3FqEXaI5ehM8PXnGXWnNLxIA7OmSzr8S59FvFuZTrLCtIBQkJqts8BykUKLssXJp91HEE9P%2bni2g60w6gU6PBc1U6p1JtmxFuSI1LHbdgnMY%2bbZqAhfHmPS1Wq84k3VKHnjuRQk5cYBJIcod11hFkKaiPiw9rH8w2%2fmkaN%2fRL2rRpuM7I6GIokPCK61Fpeiyf%2b2emcd3ooegkUTGeiwebYRXzGbpNPLyQITmpcNe5lHMtz23KLFJV%2b5ndMgfiHosGvNAe115Ej4ZLCQ1ElPUvRRkL6%2b%2bHiUnaYV0HdLxFJSNto2V8HyaY%2fcjfRZHmtG12vs7ScDVJNaiw38rp7IYkvjSB5R3M5aab04ZViHW6XAKwZ2HcYQVgrIOKqlXyQ%3d%3d&game_biz=hk4e_cn&end_id=1654229160001129743&gacha_type=301&page=16&size=6`
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func TestBili(t *testing.T) {
	client := resty.New()

	res, _ := client.R().Get("https://bilibili.com")

	if res.StatusCode() != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode(), res.Status())
	}
	fmt.Println(res.String())
}
