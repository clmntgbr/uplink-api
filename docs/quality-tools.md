# Quality Assurance Tools

This project includes three quality assurance tools that run automatically in CI/CD and can be run locally.

## Tools

### PHPStan - Static Analysis
PHPStan analyzes your code without running it to find bugs and errors.

**Run locally:**
```bash
composer phpstan
```

**Configuration:** `phpstan.neon.dist`

### PHP CS Fixer - Code Style
PHP CS Fixer automatically formats your code according to coding standards.

**Check code style:**
```bash
composer cs-check
```

**Fix code style:**
```bash
composer cs-fix
```

**Configuration:** `.php-cs-fixer.dist.php`

### Rector - Automated Refactoring
Rector automatically upgrades and refactors your code to modern PHP standards.

**Check for refactoring opportunities:**
```bash
composer rector-check
```

**Apply refactoring:**
```bash
composer rector-fix
```

**Configuration:** `rector.php`

## Run All QA Tools

To run all quality assurance tools at once:

```bash
composer qa
```

This will run PHPStan, PHP CS Fixer (check mode), and Rector (check mode).

## CI/CD Integration

All three tools run automatically in GitHub Actions CI pipeline:
- **PHPStan Job**: Runs static analysis
- **PHP CS Fixer Job**: Checks code style compliance
- **Rector Job**: Checks for refactoring opportunities

The CI jobs will fail if any tool detects issues that need to be fixed.

## Docker Usage

You can also run these tools inside the Docker container:

```bash
docker compose exec php vendor/bin/phpstan analyse
docker compose exec php vendor/bin/php-cs-fixer fix --dry-run --diff
docker compose exec php vendor/bin/rector process --dry-run
```

