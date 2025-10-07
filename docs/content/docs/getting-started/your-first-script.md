+++
title = 'Your First Script'
linkTitle = 'Your First Script'
description = 'Instructions for creating your first script using DLiteScript.'
weight = 0
draft = false
+++

If you haven't already installed DLiteScript, please consult the [installation guide](../installation).

1. Create a file called `main.dl` in any directory you like

1. Copy and paste the following contents in the file

   ```go
   var person string = "John"
   printf("Hello, %s!\n", person)
   ```

1. Run the script

   ```bash
   dlitescript main.dl
   ```

You should now see `Hello, John!` in your terminal.

That's it! You just made your first script in DLiteScript.
