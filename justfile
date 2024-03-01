set positional-arguments

@run env:
    ./script/run.sh $1

@revert-migration env:
    ./script/revert-migration.sh $1

@apply-migrations env:
    ./script/apply-migrations.sh $1

@gen-models env:
    ./script/gen-models.sh $1

@http-test env: 
    ./script/http-test.sh $1
