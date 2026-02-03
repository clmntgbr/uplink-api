<?php

declare(strict_types=1);

namespace App\Infrastructure\Project\Validation\Constraint;

use Attribute;
use Symfony\Component\Validator\Constraint;

#[Attribute(Attribute::TARGET_PROPERTY)]
final class MaxWorkflowsPerProject extends Constraint
{
    public const int MAX_WORKFLOWS = 50;
    public const string CODE = 'max.workflows.per.project';
    public const string GROUP_CREATE = 'workflow:create';

    public string $message = 'You have reached the maximum number of workflows ({{ limit }}).';
}
