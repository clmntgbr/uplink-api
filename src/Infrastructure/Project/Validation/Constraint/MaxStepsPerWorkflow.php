<?php

declare(strict_types=1);

namespace App\Infrastructure\Project\Validation\Constraint;

use Attribute;
use Symfony\Component\Validator\Constraint;

#[Attribute(Attribute::TARGET_PROPERTY)]
final class MaxStepsPerWorkflow extends Constraint
{
    public const int MAX_STEPS = 100;
    public const string CODE = 'max.steps.per.workflow';
    public const string GROUP_CREATE = 'step:create';

    public string $message = 'You have reached the maximum number of steps ({{ limit }}).';
}
