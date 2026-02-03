<?php

declare(strict_types=1);

namespace App\Infrastructure\Project\Validation\Constraint;

use Attribute;
use Symfony\Component\Validator\Constraint;

#[Attribute(Attribute::TARGET_PROPERTY)]
final class MaxProjectsPerUser extends Constraint
{
    public const int MAX_PROJECTS = 10;
    public const string CODE = 'max.projects.per.user';

    public string $message = 'You have reached the maximum number of projects ({{ limit }}).';
}
