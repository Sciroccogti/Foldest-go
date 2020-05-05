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

## Development progress
 
- [ ] Automatic
- [x] Temp trash bin
- [ ] Customize rules
