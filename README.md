# NIWT

A very simplistic weather client for the terminal based on **[wego](https://github.com/schachmat/wego)**

## Description

This is a very fast rewrite of the **wego** project, to fit my specific use-case, where eww doesn't play well with it, especially when trying to fit the data in a smaller space. It has only the bare functionality to work with my specific use-case on my specific machine.

There should be no reason for almost anyone to be using this instead of the main **wego** project or **wttr.in**.

Tyvm :heart:

## Usage

Won't include detailed usage, but basic example is:

```bash
$ NIWT Berlin
                             ┌─────────────┐
┌─────────────────┬──────────┤ Sat 26. Oct ├──────────┬─────────────────┐
│     Morning     │      Noon└──────┬──────┘Evening   │     Night       │
├─────────────────┼─────────────────┼─────────────────┼─────────────────┤
│ few clouds      │ clear sky       │ broken clouds   │ scattered cloud │
│    \__/         │     \ . /       │                 │                 │
│  __/  . 9 °C    │    - .- 15 °C   │      .- 15 °C   │      .- 12 °C   │
│    \_(  ↑ 5 km/h│   ‒ (   ↖ 3 km/h│   .-(   ↖ 8 km/h│   .-(   ↖ 9 km/h│
│    /(__ 10 km   │    . `- 10 km   │  (___._ 10 km   │  (___._ 10 km   │
│         0.0 mm/h│     / ' 0.0 mm/h│         0.0 mm/h│         0.0 mm/h│
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
```

## Contribution & Issues

Feel free to do anything you want, don't mean I'll do anything about it.

## FAQ

- Why didn't I just fork **wego**? Idk ¯\\\_(ツ)\_/¯

## TODO's

- Write proper documentation
- Add automatic location detection
- Properly handle configuration file
- Make it less chatjipity
- Just all around improvement on everything
