<?php

declare(strict_types=1);

namespace App\Shared\Infrastructure\Bus;

use App\Shared\Application\Bus\QueryBusInterface;
use Override;
use RuntimeException;
use Symfony\Component\Messenger\MessageBusInterface;
use Symfony\Component\Messenger\Stamp\HandledStamp;
use Symfony\Component\Messenger\Stamp\StampInterface;

use function sprintf;

class QueryBus implements QueryBusInterface
{
    public function __construct(private readonly MessageBusInterface $queryBus)
    {
    }

    /**
     * @param array<int, StampInterface> $stamps
     */
    #[Override]
    public function dispatch(object $query, array $stamps = []): mixed
    {
        $envelope = $this->queryBus->dispatch($query);
        $handledStamp = $envelope->last(HandledStamp::class);

        if (! $handledStamp instanceof StampInterface) {
            throw new RuntimeException(sprintf('No handler found for query of type "%s".', $query::class));
        }

        return $handledStamp->getResult();
    }
}
