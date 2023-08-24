### STAMP PDF

The application is setting watermark stamp using the png file.

### Commands

```
stamp
stamp-init
help
```

### Flags

- global: ```--file | -o, --out | --stamp-file```
- for stamp-init: ```--text | --executor```
- for stamp: ```--stamp | --text | --executor | -p, -page | --pos```

### Example commands

> creating stamp from your template.png with fill this

```bash
stamp-init --file $PWD/stamp_template.png \
-o $PWD/stamp.png \
--text "Присвоен УНВД 0/123/4567/8910 от 01.01.2023г." \
--executor "Выскочка Евгений Анатольевич (Сеньор помидор)"
```

> pre-initial stamp created

```bash
stamp --file $PWD/in.pdf \
--stamp-file $PWD/stamp.png \
-p 1- \
-o $PWD/out.pdf
```

> using a stamp template with the default name

```stamp_template.png```

```bash
stamp --file $PWD/in.pdf \
-o $PWD/out.pdf \
--stamp true \
--text "Присвоен УНВД 0/123/4567/8910 от 01.01.2023г." \
--executor "Выскочка Евгений Анатольевич (Сеньор помидор)"
```
