<?php

declare(strict_types=1);

namespace App\Shared\Application\Bus;

use Symfony\Component\Messenger\Stamp\StampInterface;

interface QueryBusInterface
{
    /**
     * @param array<int, StampInterface> $stamps
     */
    public function dispatch(object $message, array $stamps = []): mixed;
}
