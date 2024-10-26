package frontends

import (
	"NIWT/iface"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/mattn/go-runewidth"
)

func aatPad(s string, mustLen int) (ret string) {
	ansiEsc := regexp.MustCompile("\033.*?m")
	ret = s
	realLen := runewidth.StringWidth(ansiEsc.ReplaceAllLiteralString(s, ""))
	delta := mustLen - realLen
	if delta > 0 {
		ret += strings.Repeat(" ", delta)
	} else if delta < 0 {
		toks := ansiEsc.Split(s, 2)
		tokLen := runewidth.StringWidth(toks[0])
		if tokLen > mustLen {
			ret = fmt.Sprintf("%.*s", mustLen, toks[0])
		} else {
			esc := ansiEsc.FindString(s)
			ret = fmt.Sprintf("%s%s%s", toks[0], esc, aatPad(toks[1], mustLen-tokLen))
		}
	}
	return
}

func formatWind(deg int) string {
	arrows := []string{"↓", "↙", "←", "↖", "↑", "↗", "→", "↘"}
	return arrows[((deg+22)%360)/45]
}

func parseCode(cond iface.WeatherSpan) (icon []string) {

	codemap := map[int]iface.WeatherCode{
		200: iface.CodeThunderStorm,
		201: iface.CodeThunderStorm,
		210: iface.CodeThunderStorm,
		230: iface.CodeThunderStorm,
		231: iface.CodeThunderStorm,
		202: iface.CodeThunderStorm,
		211: iface.CodeThunderStorm,
		212: iface.CodeThunderStorm,
		221: iface.CodeThunderStorm,
		232: iface.CodeThunderStorm,
		300: iface.CodeRainShower,
		301: iface.CodeRainShower,
		310: iface.CodeRainShower,
		311: iface.CodeRainShower,
		313: iface.CodeRainShower,
		321: iface.CodeRainShower,
		302: iface.CodeRainShower,
		312: iface.CodeRainShower,
		314: iface.CodeRainShower,
		500: iface.CodeRain,
		501: iface.CodeRain,
		502: iface.CodeRain,
		503: iface.CodeRain,
		504: iface.CodeRain,
		511: iface.CodeSnow,
		520: iface.CodeRainShower,
		521: iface.CodeRainShower,
		522: iface.CodeRainShower,
		531: iface.CodeRainShower,
		600: iface.CodeSnow,
		601: iface.CodeSnow,
		602: iface.CodeSnow,
		611: iface.CodeSnow,
		612: iface.CodeSnow,
		615: iface.CodeSnow,
		616: iface.CodeSnow,
		620: iface.CodeSnow,
		621: iface.CodeSnow,
		622: iface.CodeSnow,
		701: iface.CodeMist,
		711: iface.CodeMist,
		721: iface.CodeMist,
		741: iface.CodeMist,
		731: iface.CodeMist, // sand, dust whirls
		751: iface.CodeMist, // sand
		761: iface.CodeMist, // dust
		762: iface.CodeMist, // volcanic ash
		771: iface.CodeMist, // squalls
		781: iface.CodeMist, // tornado
		800: iface.CodeClear,
		801: iface.CodeFewClouds,
		802: iface.CodeCloudy,
		803: iface.CodeCloudy,
		804: iface.CodeCloudy,
	}

	codes := map[iface.WeatherCode][]string{
		iface.CodeUnknown: {
			"    .-.      ",
			"     __)     ",
			"    (        ",
			"     `-᾿     ",
			"      •      ",
		},
		iface.CodeCloudy: {
			"             ",
			"     .--.    ",
			"  .-(    ).  ",
			" (___.__)__) ",
			"             ",
		},
		iface.CodeMist: {
			"             ",
			" _ - _ - _ - ",
			"  _ - _ - _  ",
			" _ - _ - _ - ",
			"             ",
		},
		iface.CodeRain: {
			"     .-.     ",
			"    (   ).   ",
			"   (___(__)  ",
			"    ´ ´ ´ ´  ",
			"   ´ ´ ´ ´   ",
		},
		iface.CodeRainShower: {
			" _`/\"\".-.  ",
			"  ,\\_(   ). ",
			"   /(___(__) ",
			"     ´ ´ ´ ´ ",
			"    ´ ´ ´ ´  ",
		},
		iface.CodeSnow: {
			"     .-.     ",
			"    (   ).   ",
			"   (___(__)  ",
			"    *  *  *  ",
			"   *  *  *   ",
		},
		iface.CodeFewClouds: {
			"   \\__/     ",
			" __/  .-.    ",
			"   \\_(   ). ",
			"   /(___(__) ",
			"             ",
		},
		iface.CodeClear: {
			"    \\ . /   ",
			"   - .-. -   ",
			"  ‒ (   ) ‒  ",
			"   . `-᾿ .   ",
			"    / ' \\   ",
		},
		iface.CodeThunderStorm: {
			"     .-.     ",
			"    (   ).   ",
			"   (___(__)  ",
			"  ´_\\ ´´/_  ",
			"  ´\\ ´ ´ /  ",
		},
	}

	if val, ok := codemap[cond.Weather[0].ID]; ok {
		icon = codes[val]
	}

	return

}

func formatCond(cur []string, cond iface.WeatherSpan, current bool) (ret []string) {

	textWidth := 8
	iconWidth := 7

	icon := parseCode(cond)

	ret = append(ret, fmt.Sprintf("%v %v ", cur[0], aatPad(cond.Weather[0].Description, textWidth+iconWidth)))
	ret = append(ret, fmt.Sprintf("%v %v ", cur[1], aatPad(icon[0], iconWidth+textWidth)))
	ret = append(ret, fmt.Sprintf("%v %v %v", cur[2], aatPad(icon[1], iconWidth), aatPad(fmt.Sprintf("%.0f °C", cond.Main.TempC), textWidth)))
	ret = append(ret, fmt.Sprintf("%v %v %v", cur[3], aatPad(icon[2], iconWidth), aatPad(fmt.Sprintf("%s %.0f km/h", formatWind(cond.Wind.Deg), cond.Wind.Speed*3.6), textWidth)))
	ret = append(ret, fmt.Sprintf("%v %v %v", cur[4], aatPad(icon[3], iconWidth), aatPad(fmt.Sprintf("%d km", cond.Visibility/1000), textWidth)))
	ret = append(ret, fmt.Sprintf("%v %v %v", cur[5], aatPad(icon[4], iconWidth), aatPad(fmt.Sprintf("%.1f mm/h", cond.Rain.MM3h/3), textWidth)))

	return
}

func printDay(currentDay []iface.WeatherSpan) (ret []string) {

	timeSlots := []iface.WeatherSpan{
		currentDay[0],
		currentDay[2],
		currentDay[4],
		currentDay[6],
	}

	ret = make([]string, 6)
	for i := range ret {
		ret[i] = "│"
	}

	for _, s := range timeSlots {
		ret = formatCond(ret, s, false)
		for i := range ret {
			ret[i] = ret[i] + "│"
		}
	}

	dateFmt := "┤ " + time.Unix(currentDay[0].Dt, 0).Format("Mon 02. Jan") + " ├"

	ret = append([]string{
		"                             ┌─────────────┐                             ",
		"┌─────────────────┬──────────" + dateFmt + "──────────┬─────────────────┐",
		"│     Morning     │      Noon└──────┬──────┘Evening   │     Night       │",
		"├─────────────────┼─────────────────┼─────────────────┼─────────────────┤"},
		ret...)
	ret = append(ret,
		"└─────────────────┴─────────────────┴─────────────────┴─────────────────┘")

	for _, line := range ret {
		fmt.Println(line)
	}
	return
}

func AsciiDraw(weather iface.WeatherResponse) {

	printDay(weather.List)
}
