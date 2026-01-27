<?php

declare(strict_types=1);

namespace App\Shared\Domain\Repository;

use Doctrine\Bundle\DoctrineBundle\Repository\ServiceEntityRepository;
use Symfony\Component\Uid\Uuid;

/**
 * @template T of object
 *
 * @extends ServiceEntityRepository<T>
 */
abstract class AbstractRepository extends ServiceEntityRepository
{
    /**
     * @param T $entity
     *
     * @return T
     */
    public function refresh(object $entity): object
    {
        $this->getEntityManager()->refresh($entity);

        return $entity;
    }

    /**
     * @return T|null
     */
    public function findByUuid(Uuid $id): ?object
    {
        return $this->findOneBy(['id' => $id]);
    }

    /**
     * @param T $entity
     */
    public function save(object $entity, bool $force = false): void
    {
        if ($this->getEntityManager()->contains($entity)) {
            $this->flush();
        }

        $identifier = $this->getIdentifier($entity);

        if ($identifier === null) {
            $this->getEntityManager()->persist($entity);
        }

        if ($force) {
            $this->getEntityManager()->persist($entity);
        }

        $this->flush();
    }

    /**
     * @param T $entity
     */
    public function delete(object $entity): void
    {
        $this->getEntityManager()->remove($entity);
    }

    public function flush(): void
    {
        $this->getEntityManager()->flush();
    }

    /**
     * @param list<T> $entities
     */
    public function saveAll(array $entities): void
    {
        foreach ($entities as $entity) {
            $this->save($entity);
        }
    }

    /**
     * @param T $entity
     */
    private function getIdentifier(object $entity): ?string
    {
        if (method_exists($entity, 'getId')) {
            return (string) $entity->getId();
        }

        return null;
    }
}
