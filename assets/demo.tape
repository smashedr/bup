Output dist/demo.gif
#Output dist/demo.mp4
#Output dist/demo.webm

Require bup

Set Theme "Catppuccin Frappe"
Set MarginFill "#3590D6"
Set Margin 10
Set WindowBar Colorful
Set BorderRadius 12
Set Padding 20
Set Width 1200
Set Height 900
Set FontSize 34

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
