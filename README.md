all two components function are work  testing result below
-----------------------------------------------------------
1. for get all folders function:
(base) PS C:\Users\28068\Desktop\sc-take-home-intern-2024_2025\sc-take-home-assessment-take-home-2025> go test ./folder/get_folder_test.go -v 
go: downloading github.com/gofrs/uuid v4.3.0+incompatible
go: downloading github.com/lucasepe/codename v0.2.0
=== RUN   Test_folder_GetAllChildFolders
=== RUN   Test_folder_GetAllChildFolders/Get_child_folders_for_alpha
=== RUN   Test_folder_GetAllChildFolders/Get_child_folders_for_bravo
=== RUN   Test_folder_GetAllChildFolders/No_child_folders_for_echo
=== RUN   Test_folder_GetAllChildFolders/No_child_folders_for_echo#01
=== RUN   Test_folder_GetAllChildFolders/No_child_folders_for_org2
--- PASS: Test_folder_GetAllChildFolders (0.00s)
    --- PASS: Test_folder_GetAllChildFolders/Get_child_folders_for_alpha (0.00s)
    --- PASS: Test_folder_GetAllChildFolders/Get_child_folders_for_bravo (0.00s)
    --- PASS: Test_folder_GetAllChildFolders/No_child_folders_for_echo (0.00s)
    --- PASS: Test_folder_GetAllChildFolders/No_child_folders_for_echo#01 (0.00s)
    --- PASS: Test_folder_GetAllChildFolders/No_child_folders_for_org2 (0.00s)
PASS
ok      command-line-arguments  0.398s


-------------------------------------------------------------
2. for move folders function:
(base) PS C:\Users\28068\Desktop\sc-take-home-intern-2024_2025\sc-take-home-assessment-take-home-2025> go test ./folder/move_folder_test.go -v
go: downloading github.com/stretchr/testify v1.9.0
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading gopkg.in/yaml.v3 v3.0.1
go: downloading github.com/pmezard/go-difflib v1.0.0
=== RUN   Test_folder_MoveFolder
=== RUN   Test_folder_MoveFolder/Move_bravo_to_delta
=== RUN   Test_folder_MoveFolder/Move_bravo_to_itself
=== RUN   Test_folder_MoveFolder/Move_bravo_to_charlie
=== RUN   Test_folder_MoveFolder/Move_bravo_to_foxtrot
--- PASS: Test_folder_MoveFolder (0.00s)
    --- PASS: Test_folder_MoveFolder/Move_bravo_to_delta (0.00s)
    --- PASS: Test_folder_MoveFolder/Move_bravo_to_itself (0.00s)
    --- PASS: Test_folder_MoveFolder/Move_bravo_to_charlie (0.00s)
    --- PASS: Test_folder_MoveFolder/Move_bravo_to_foxtrot (0.00s)
PASS
ok      command-line-arguments  0.467s
