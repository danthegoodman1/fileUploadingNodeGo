Compiled GO Server:

```
    time_namelookup:  0.004323s
       time_connect:  0.004512s
    time_appconnect:  0.000000s
   time_pretransfer:  0.004559s
      time_redirect:  0.000000s
 time_starttransfer:  0.004566s
                    ----------
         time_total:  0.005046s
```


NodeJS Server:

```
    time_namelookup:  0.004331s
       time_connect:  0.004529s
    time_appconnect:  0.000000s
   time_pretransfer:  0.004565s
      time_redirect:  0.000000s
 time_starttransfer:  0.004571s
                    ----------
         time_total:  0.007928s
```

Some `time_total`s were seen > `0.01` for NodeJS
