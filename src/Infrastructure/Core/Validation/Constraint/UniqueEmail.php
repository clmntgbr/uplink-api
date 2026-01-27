<?php

declare(strict_types=1);

namespace App\Infrastructure\Core\Validation\Constraint;

use Attribute;
use Override;
use Symfony\Component\Validator\Constraint;

#[Attribute]
class UniqueEmail extends Constraint
{
    public string $message = 'This email address is already registered';

    #[Override]
    public function validatedBy(): string
    {
        return static::class . 'Validator';
    }
}
