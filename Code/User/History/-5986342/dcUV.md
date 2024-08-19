# GIT ![Static Badge](https://img.shields.io/badge/git-2.45.2-red)

This project is designed to introduce us to the world of version control and collaboration using Git. Git is a powerful and widely used tool for tracking changes to our projects, collaborating with others and ensuring the integrity of our code.

Throughout this project, we'll embark on a journey that will enable us to gradually develop our Git skills. Starting with the basics, we'll gradually explore more advanced topics, equipping ourselves with the knowledge and practices essential for effective version control and collaboration.

Get ready for Git!

## Setting Up Git

- Install Git: (linux)

  ```bash
  sudo apt install git-all
  ```

- Configure Git:

  ```bash
  git config --global user.name "John Smith"
  git config --global user.email "johnsmith@exemple.com"
  ```

## Git commits to commit

- Folder, subfolder and file creation:

  ```bash
  $ pwd
  /home/username/Documents/Zone01/git

  $ mkdir work
  $ cd work
  $ mkdir hello
  $ cd hello
  $ echo 'echo "Hello, World"' > hello.sh
  ```

- Initialize the git repository in the hello directory:

  ```
  git init
  ```

- Check the status and act accordingly with the output of the executed command:

  ```bash
  $ git status # check status
  Sur la branche master

  Aucun commit

  Fichiers non suivis:
  (utilisez "git add <fichier>..." pour inclure dans ce qui sera validé)
          hello.sh

  aucune modification ajoutée à la validation mais des fichiers non suivis sont présents (utilisez "git add" pour les suivre)
  $ git add hello.sh # stage hello.sh
  ```

- Change file contents as requested

- Stage the changed file and commit the changes, the working tree should be clean:

  ```bash
  $ git add hello.sh
  $ git commit -m "first commit: hello.sh" # commit all staged files

  $ git status
  Sur la branche master
  rien à valider, la copie de travail est propre
  ```

- Modify the file by manking two different commits: (add a comment on line 3 + add/modify lines 4 and 5)

  ```bash
  # replaces the entire contents of the file with the requested content

  $ git add -p # then e + enter
  # delete the last two lines* & exit by saving
  $ git commit -m "comment on line 3" # commit 1/2

  $ git add hello.sh # stage with lines 4 and 5
  $ git commit -m "add/modify lines 4 and 5"
  ```

  ##### _The "<ins>git add -p</ins>" command goes through each modified file in our working directory and presents each modification differently. For each modification, it asks whether we wish to include it in the commit, ignore it or open the file in a text editor for further review._

  ##### \*the two lines we don't want to keep for this first of the two commits

## History

- history commands :

  ```bash
  $ git log # show the history
  ...
  $ git log --oneline # one-line history
  46c65a0 (HEAD -> master) add/modify lines 4 and 5
  6f3f366 comment on line 3
  0d3cb83 first commit: hello.sh
  ```

### Controlled Entries:

- show last 2 entries:

  ```
  git log -n 2
  ```

- show commits made within the last 5 minutes:

  ```
  git log --since=5.minutes
  ```

### Personalized Format:

- Show logs in a personalized format: (commit hash, date, message, branch information, and author name)

  ```
  git log --pretty=format:"%h %as | %s %d [%an]"
  ```

## Check it out

### Restore First Snapshot:

- Revert the working tree to its initial state, as captured in the first snapshot, and then print the content of the hello.sh file:

  ```
  $ git checkout HEAD~^
  # or
  $ git checkout [first commit hash]

  $ cat hello.sh
  #!/bin/bash

  echo "Hello, $1"
  ```

  ##### _<ins>git checkout HEAD~^</ins> = Checkout the parent of the first commit_

### Restore Second Recent Snapshot:

- Revert the working tree to the second most recent snapshot and print the content of the hello.sh file:

  ```
  $ git checkout [second most recent commit hash]
  $ cat hello.sh
  #!/bin/bash

  # Default is "World"
  ```

### Return to Latest Version:

- Ensure that the working directory reflects the latest version of the hello.sh file present in the main branch, without referring to specific commit hashes:

  ```
  git checkout master
  ```

## TAG me

### Referencing Current Version:

- Tag the current version of the repository as v1:

  ```bash
  git tag v1 # tag the current commit
  ```

### Tagging Previous Version:

- Tag the version immediately prior to the current version as v1-beta, without relying on commit hashes to navigate through the history:

  ```bash
  git tag v1-beta HEAD^ # ^ for HEAD-1
  ```

### Navigating Tagged Versions:

- Move back and forth between the two tagged versions, v1 and v1-beta:

  ```bash
  $ git checkout v1-beta
  # and
  $ git checkout v1
  ```

### Listing Tags:

- Display a list of all tags present in the repository to verify successful tagging:

  ```
  $ git tag -l
  v1
  v1-beta
  ```

## Changed your mind?

### Reverting Changes:

- Modify the latest version of the file with unwanted comments, then revert it back to its original state before staging using a Git command:

  ```
  git restore <file_name>
  ```

### Staging and Cleaning:

- Introduce unwanted changes to the file, stage them, then clean the staging area to discard the changes:

  ```
  $ git add <file_name>
  $ git restore --staged <file_name>
  ```

### Committing and Reverting:

- Add the following unwanted changes again, stage the file, commit the changes, then revert them back to their original state:

  ```bash
  $ git add <file_name>
  $ git commit -m "Committing and Reverting:"

  $ git revert HEAD
  # or
  $ git revert <commit_hash>
  ```

### Tagging and Removing Commits:

- Tag the latest commit with oops, then remove commits made after the v1 version. Ensure that the HEAD points to v1:

  ```bash
  git tag oops
  git reset (--hard) v1 # --hard en fonction de la situation, ici ne change rien
  ```

  ##### _"<ins>git reset v1</ins>" deletes all commits after v1_

### Displaying Logs with Deleted Commits:

- Show the logs with the deleted commits displayed, particularly focusing on the commit tagged oops:

  ```
  git log --all
  ```

### Cleaning Unreferenced Commits:

- Ensure that unreferenced commits are deleted from the history, meaning there should be no logs for these deleted commits:

  ```
  git tag -d oops
  ```

  ##### _in our case, deleting the "oops" tag automatically deletes its associated commit and its parent_

### Author Information:

- Add an author comment to the file and commit the changes:

  ```
  git add hello.sh
  git commit -m "add author comment"
  ```

- Oops the author email was forgotten, update the file to include the email without making a new commit, but include the change in the last commit:

  ```bash
  # add author email, then
  $ git add hello.sh
  $ git commit --amend -C HEAD
  ```

  ##### _"<ins>--amend</ins>" to modify the last commit instead of creating a new one_

  ##### _"<ins>-C HEAD</ins>" to use the same commit message as the last commit_

## Move it

### Moving hello.sh:

- Using Git commands, move the program hello.sh into a lib/ directory, and then commit the move:

  ```
  mkdir lib
  git mv hello.sh lib/
  ```

- Create a Makefile in the root directory of the repository with the provided content and commit it to the repository:

  ```
  touch MakeFile
  # then add the requested content to it
  ```

## blobs, trees and commits

### Exploring .git/ Directory:

- Navigate to the .git/ directory in your project and examine its contents.You will have to explain the purpose of each subdirectory, including objects/, config, refs, and HEAD in the audit:

  1. <ins>objects/</ins>

     This directory stores the actual content of your project in a compressed and deduplicated format. It contains various subdirectories and files representing different object types:

     - <ins>blobs</ins>: store the raw content of tracked files in your project
     - <ins>trees</ins>: represent a directory structure by referencing the blobs for its files and subdirectories
     - <ins>commits</ins>: store the metadata for each commit, including the commit message, author information, and references to the parent tree(s)

  2. <ins>config</ins>

     This file stores Git configuration settings specific to your repository. It defines parameters like username, email, preferred editor, and more.

  3. <ins>refs/</ins>

     This directory stores references to various Git objects, most importantly commits. It contains several subdirectories:

     - <ins>heads</ins>: stores references (branch names) that point to specific commits. The file named HEAD within this subdirectory points to the currently active branch
     - <ins>tags</ins>: stores references (tag names) that also point to specific commits. Tags are used to mark specific versions of your project
     - <ins>remotes</ins>: (Optional) can store references to remote repositories for collaboration purposes

  4. <ins>HEAD</ins>

     This is a symbolic link (not a directory) located at the root of the .git directory. It points to the currently active branch, which itself is a reference to a specific commit in the refs/heads subdirectory.

### Latest Object Hash:

- Find the latest object hash within the .git/objects/ directory using Git commands and print the type and content of this object using Git commands:

  ```
  $ git rev-parse HEAD
  <commit_hash>
  $ git cat-file -t <commit_hash>
  commit
  $ git cat-file -p <commit_hash>
  tree b7b49c6ea13e1e1423ca89787cbc1882f8eb0eac
  parent 4fcffcf2962157ab621848ba4ea85e9895b9ad73
  author Adam GONCALVES <goncalvesadam@icloud.com> 1720448920 +0200
  committer Adam GONCALVES <goncalvesadam@icloud.com> 1720448920 +0200

  add MakeFile
  $
  ```

### Dumping Directory Tree:

- Use Git commands to dump the directory tree referenced by this commit:

  ```
  git ls-tree -r <commit_hash>
  ```

- Dump the contents of the lib/ directory and the hello.sh file using Git commands:

  ```
  $ git ls-tree HEAD lib/
  100644 blob <blob_hash>    lib/hello.sh
  $ git cat-file -p <blob_hash>
  ```

## Branching

### Create and Switch to New Branch:

- Create a local branch named greet and switch to it:

  ```
  git branch greet
  git checkout greet
  ```

- In the lib directory, create a new file named greeter.sh and add the provided code to it. Commit these changes:

  ```bash
  $ touch lib/greeter.sh
  # add the requested content to it
  $ git add lib/greeter.sh
  $ git commit -m "add lib/greeter.sh"
  ```

- Update the lib/hello.sh file by adding the requested content, stage and commit the changes

- Update the Makefile with the requested content and commit the changes

- Switch back to the main branch, compare and show the differences between the main and greet branches for Makefile, hello.sh, and greeter.sh files:

  ```
  git checkout master
  git diff master greeter MakeFile
  git diff master greeter lib/hello.sh
  ```

- Generate a README.md file for the project with the requested content. Commit this file

- Draw a commit tree diagram illustrating the diverging changes between all branches to demonstrate the branch history:

  ```
  $ git log --graph --decorate --oneline --all
  * 7bd25fd (HEAD -> master) add README
  | * f77de7a (greeter) add comment to MakeFile
  | * f69e696 update hello.sh
  | * d73cc71 add lib/greeter.sh
  |/
  * fe361b3 add MakeFile
  * 4fcffcf move hello.sh to lib/
  * b62a717 add author comment
  * 46c65a0 (tag: v1) add/modify lines 4 and 5
  * 6f3f366 (tag: v1-beta) comment on line 3
  * 0d3cb83 first commit: hello.sh
  $
  ```

## Conflicts, merging and rebasing

### Merge Main into Greet Branch:

- Merging the changes from the main branch into the greet branch:

  ```bash
  $ git checkout greet
  $ git merge master
  # save and exit
  ```

- Switch to main branch and make the changes requested to the hello.sh file, save and commit the changes

### Merging Main into Greet Branch (Conflict):

- Attempt to merge the main branch into greet. Bingooo! There you have it, a conflict

- Resolve the conflict (manually or using graphical merge tools), accept changes from main branch, then commit the conflict resolution:

  1. open the file concerned
  2. delete the unwanted part and the markers added by git
  3. save the file
  4. git add and git commit

### Rebasing Greet Branch:

- Go back to the point before the initial merge between main and greet:

  ```
  git checkout <commit_hash>
  ```

- Rebase the greet branch on top of the latest changes in the main branch:

  ```
  git rebase master
  ```

### Merging Greet into Main:

- Merge the changes from the greet branch into the main branch:

  ```
  git checkout master
  git merge greet
  ```

### Understanding Fast-Forwarding and Differences:

- Explain fast-forwarding and the difference between merging and rebasing:

  1. Fast-forwarding:

     - This is a specific type of merge that occurs when the main branch can be directly integrated with the greet branch by simply moving the HEAD pointer on main forward
     - It happens when the history of the greet branch is a linear continuation of the main branch (no divergences)
     - No new merge commit is created in this case
     - The git merge --ff-only command attempts a fast-forward merge, but it will fail if there are any conflicts

  2. Merging:

     - Merging is a more general approach for combining changes from two branches
     - Git creates a new merge commit that references the two heads (commits) being merged
     - This commit serves as a record of the merge operation and helps track the history of both branches
     - Merging can handle conflicts that might arise when changes have been made to the same files in both branches
     - Use git merge <branch_name> to initiate a merge

  3. Rebasing:

     - Rebasing is a technique for rewriting the history of a branch
     - It replays the commits of a branch (e.g., greet) on top of another branch (e.g., main)
     - Git essentially detaches the commits from their original position and reapplies them one by one
     - This creates a linear history where the greet branch appears to have directly originated from the latest commit in main
     - Rebasing can be useful if you want a clean and linear branch history, especially for development branches that haven't been shared with others
     - However, rebasing rewrites history, which can cause issues if the branch has already been shared and others have pulled or pushed their local copies
     - Use git rebase <branch_name> to initiate a rebase

  - Use <ins>Fast-forwarding</ins> only if you're certain the greet branch doesn't introduce conflicts and its history is a linear continuation of main
  - <ins>Merging</ins> is the safer and more general approach for integrating changes, especially when dealing with potential conflicts or a desire to preserve the complete branch history
  - Consider <ins>Rebasing</ins> if you want a clean linear history for your branch, but use it with caution if the branch has been shared with collaborators.

## Local and remote repositories (save here)

- In the work/ directory, make a clone of the repository hello as cloned_hello: (Do not use copy command)

  ```bash
  $ cd ..
  $ pwd
  # /home/username/Documents/Zone01/git/work

  $ git clone hello/ cloned_hello/
  ```

- Show the logs for the cloned repository:

  ```
  $ cd cloned_hello
  $ git log --oneline
  7bdc4e3 (HEAD -> master, origin/master, origin/greet, origin/HEAD) Resolved merge conflict in hello.sh
  ...
  $
  ```

- Display the name of the remote repository and provide more information about it:

  ```bash
  $ git remote -v # -v for more verbose
  origin  /home/shasor/Documents/Zone01/git/work/hello (fetch)
  origin  /home/shasor/Documents/Zone01/git/work/hello (push)
  $
  ```

  "origin" is the name of the remote, what's written just to the right of it are the Internet links, or in our case, the local access paths to our remote.

- List all remote and local branches in the cloned_hello repository:

  ```
  $ git branch -a
  * master
  remotes/origin/HEAD -> origin/master
  remotes/origin/greet
  remotes/origin/master
  $
  ```

- Make changes to the original repository, update the README.md file with the requested content, and commit (80834b1) the changes

- Inside the cloned repository (cloned_hello), fetch the changes from the remote repository and display the logs. Ensure that commits from the hello repository are included in the logs

  ```bash
  # go to cloned_hello folder then
  $ git fetch
  $ git log --oneline --all
  80834b1 (origin/master, origin/HEAD) Updated README.md with new content
  ...
  $
  ```

  it's OK, commit 80834b1 is present

- Merge the changes from the remote main branch into the local main branch:

  ```bash
  git checkout master
  git merge origin/master
  git log --oneline # to check if it's ok
  ```

- Add a local branch named greet tracking the remote origin/greet branch:

  ```bash
  $ git branch --track greet origin/greet
  la branche 'greet' est paramétrée pour suivre 'origin/greet'. # ok
  $
  ```

- Add a remote to your Git repository and push the main and greet branches to the remote:

  ```
  git remote add myremote https://github.com/username/git-remote.git # or Gitea URL
  git push myremote master greet
  ```

- "What is the single git command equivalent to what you did before to bring changes from remote to local main branch?"

  ```bash
  git pull # fetch and merge in same time
  ```

## Bare repositories
