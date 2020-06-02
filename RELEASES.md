# Releases

A document containing some important information about the major changes I made and how to migrate over.

## Version 2

### Changes
1) Reorganized and rewrote the internals of the server to give it a cleaner architecture.
2) Some changes to api endpoints. The effects are generally the same.
3) CLI tool has been updated to follow this new structure.
4) Web frontend tool has also been updated to use new api.

### Bugs

The Web UI is definitely buggy. But since you can register accounts, 
upload (and overwrite existing) packages with the CLI just fine, I'll
postpone updating the bugs till I have more time.

### How to migrate

Just run the new `docker-compose.yaml` in the `_examples` folder with the same volume. Fingers-crossed, 
it *should* be  fine. :|

The SQL script updates the database accordingly. For the server's config, check out the `config.yaml` file.
In general, the configuration has been simplified.
