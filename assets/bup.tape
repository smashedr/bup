Output dist/bup.gif
Output dist/bup.mp4
Output dist/bup.webm

Require bup

Set Theme "Catppuccin Frappe"
Set MarginFill "#6B50FF"
Set Margin 10
Set WindowBar Colorful
Set BorderRadius 10
Set Padding 30
Set Width 1200
Set Height 900
Set FontSize 26

Type "bup backup ."
Sleep 0.5s
Enter
Wait /Proceed\?/
Sleep 2s
Type "y"
Sleep 1s
Enter
Wait
Sleep 3s

Type "bup list"
Sleep 0.5s
Enter
Wait
Screenshot dist/bup.png
Sleep 2s

Type "bup l bup"
Sleep 0.5s
Enter
Wait
Sleep 2s

Type "clear"
Sleep 0.3s
Enter
