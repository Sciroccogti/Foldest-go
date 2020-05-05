# Foldest-go

Automatically manage your folder.

## How to use

1. Download binary file in release page
2. Download `conf.yml` and place it in the same folder with binary file
3. Configure settings in `conf.yml`:
```yml
verbose: # verbose output, false as default
targetdir: # the folder you want to manage, you can set it within the program. Last setted folder will be remembered.
tmpbin:
  enable: # whether to use tmpbin, false as default
  name: tmpbin/ # name of tmpbin, "tmpbin/" as default
  treshday: 30 # files not modified for more than this long will be moved into tmpbin, 30 days as default
  deleteday: 30 # files in tmpbin for more than this long will be deleted, 30 days as default
```
4. Set your rules in `rules.yml`: (Currently support 10 rules utmost)
```yml
rule1:
  enable: true
  name: document
  regex:
  - ".*?.doc"
  - ".*?.docx"
  - ".*?.pdf"
  threshday: 7
  maxsize: # MB
  minsize: # MB
rule2:
...
```

## Development progress
 
- [ ] Automatic
- [x] Temp trash bin
- [x] Customize rules
