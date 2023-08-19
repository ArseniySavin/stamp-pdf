### STAMP PDF

The application is setting watermark stamp using the png file.

### Commands

```bash
stamp --file $PWD/in.pdf \
--stamp-file $PWD/stamp.png \
-p 1- \
-o $PWD/out.pdf
```

```bash
stamp --file $PWD/in.pdf \
-o $PWD/out.pdf \
--stamp true \
--stamp-text "Присвоен УНВД 0/123/4567/8910 от 01.01.2023г." \
--stamp-executor "Выскочка Евгений Анатольевич (Сеньор помидор)"
```

```bash
stamp-init --file $PWD/stamp_template.png \
-o $PWD/stamp.png \
--stamp-text "Присвоен УНВД 0/123/4567/8910 от 01.01.2023г." \
--stamp-executor "Выскочка Евгений Анатольевич (Сеньор помидор)"
```
