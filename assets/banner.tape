Output dist/banner.gif
#Output dist/banner.mp4
#Output dist/banner.webm

Require bup

Set Theme "Catppuccin Frappe"
#Set MarginFill "#6B50FF"
#Set Margin 10
#Set WindowBar Colorful
#Set BorderRadius 10
Set Padding 84
Set Width 1280
Set Height 640
Set FontSize 24

Hide
#Type "bup --help | head -n 16 && echo '  ...'"
Type "bup --help"
Enter
Show
Sleep 1.5s

Type "bup backup ."
Sleep 0.5s
Enter
Wait /Yes./
Sleep 1s
Right
Sleep 0.4s
Left
Sleep 0.4s

Enter
Sleep 3s

Type "bup list"
Sleep 0.5s
Enter
Wait
#Screenshot dist/bup.png
Sleep 2s

Type "bup l bup"
Sleep 0.5s
Enter
Wait
Sleep 2s

Type "clear"
Sleep 0.3s
Enter
