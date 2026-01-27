<?php

declare(strict_types=1);

namespace App\Shared\Infrastructure\Bus;

use App\Shared\Application\Bus\CoreBusInterface;
use App\Shared\Application\Command\AsynchronousCoreInterface;
use Override;
use RuntimeException;
use Symfony\Component\Messenger\Envelope;
use Symfony\Component\Messenger\MessageBusInterface;

class CoreBus implements CoreBusInterface
{
    public function __construct(
        private readonly MessageBusInterface $coreBus,
    ) {
    }

    #[Override]
    public function dispatch(object $message): Envelope
    {
        if (! $message instanceof AsynchronousCoreInterface) {
            throw new RuntimeException('The message must implement AsynchronousCoreInterface.');
        }

        return $this->coreBus->dispatch($message, $message->getStamps());
    }
}
