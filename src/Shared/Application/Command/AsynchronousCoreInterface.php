<?php

declare(strict_types=1);

namespace App\Shared\Application\Command;

use JsonSerializable;
use Symfony\Component\Messenger\Stamp\StampInterface;

interface AsynchronousCoreInterface extends JsonSerializable
{
    /**
     * @return array<int, StampInterface>
     */
    public function getStamps(): array;

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array;

    public function getWebhookUrlSuccess(): string;

    public function getWebhookUrlFailure(): string;
}
