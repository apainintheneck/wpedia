# wpedia
(Built with go 1.23.2)

A very simple TUI Wikipedia client that provides the ability to search for articles and read the text portion. The parsing is not perfect by a long shot so things like tables and some lists don't show up correctly. Overall it's decent and quicker than opening a new browser window.

## Usage

```console
$ ./wpedia -h
Usage of ./wpedia:
  -l string
        2 digit language code (default "en")
```

## Examples

### Search

```console
$ wpedia warriors
```

```
    Search: warriors

    1. Warrior
  > 2. Golden State Warriors
    3. The Warriors (film)
    4. Warriors (Lin-Manuel Miranda and Eisa Davis album)
    5. The Warriors
    6. Terracotta Army
    7. Road Warriors
    8. Rainbow Warrior
    9. Hyrule Warriors
    10. Warriors (novel series)
    11. Cage Warriors
    12. Warrior (disambiguation)
    13. Samurai Warriors
    14. Dynasty Warriors
    15. Road warrior
    16. Weekend warrior

    ••••

    ↑/k up • ↓/j down • q quit • ? more

```

```
╭───────────────────────╮
│ Golden State Warriors ├─────────────────────────────────────────────────────────────
╰───────────────────────╯
    === 1946–1962: Early years in Philadelphia ===

    The Warriors were founded in 1946 as the Philadelphia Warriors, a charter member
of the Basketball Association of America. They were owned by Peter A. Tyrrell, who
also owned the Philadelphia Rockets of the American Hockey League. Tyrrell hired Eddie
Gottlieb, a longtime basketball promoter in the Philadelphia area, as coach and
general manager. The owners named the team after the Philadelphia Warriors, an old
basketball team who played in the American Basketball League in 1925.

    Led by early scoring sensation Joe Fulks, the team won the championship in the
league's inaugural 1946–47 season by defeating the Chicago Stags, four games to one.
The NBA, which was created by a 1949 merger, officially recognizes that as its own
first championship. Gottlieb bought the team in 1951.

    The Warriors won its next championship in Philadelphia in the 1955–56 season,
defeating the Fort Wayne Pistons four games to one. The Warrior stars of this era were
future Hall of Famers Paul Arizin, Tom Gola and Neil Johnston.
                                                                              ╭──────╮
──────────────────────────────────────────────────────────────────────────────┤   3% │
                                                                              ╰──────╯
```

### Random

```console
$ wpedia
```

```
    Search: Random

    1. Herbert Clifford Serasinghe
    2. Hamilton Island Cup
    3. Nosferattus palmatus
  > 4. Gamble Field
    5. 2006 Supercopa de España de Baloncesto
    6. Zemgus Girgensons
    7. Point of Rocks station
    8. Curse of Coogan's Bluff
    9. Aydar Teregulov
    10. Manny Mori
    11. Leptosia medusa
    12. Kewa Pueblo station
    13. Żmijewo-Gaje
    14. Media in Fredericton
    15. Andheri Assembly constituency
    16. Battle of Poimanenon

    ••••

    ↑/k up • ↓/j down • q quit • ? more

```

```
╭──────────────╮
│ Gamble Field ├──────────────────────────────────────────────────────────────────────
╰──────────────╯

Gamble Field was an outdoor sports stadium in the western United States, located on
the campus of the University of Colorado in Boulder. It was the predecessor of Folsom
Field.

  == History ==

  Opened 123 years ago in 1901 on September 21, it was built  via the efforts of  the
university's student body.  The field was named after Judge Harry P. Gamble, a six-
time (1891–96) football letterman and two-time captain.

  Seating capacity was initially limited to 1,000 via a 160-foot (50 m) wooden
grandstand located on the western side of the field, the only side that had seating.
The elevation of the playing field was just over 5,400 feet (1,645 m) above sea level.

  == Usage ==

                                                                              ╭──────╮
──────────────────────────────────────────────────────────────────────────────┤   0% │
                                                                              ╰──────╯
```

## Libraries

- [go-wiki](https://pkg.go.dev/github.com/trietmn/go-wiki)
- [bubbletea](https://pkg.go.dev/github.com/charmbracelet/wish/bubbletea)
