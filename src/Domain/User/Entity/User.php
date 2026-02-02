<?php

declare(strict_types=1);

namespace App\Domain\User\Entity;

use ApiPlatform\Metadata\ApiResource;
use ApiPlatform\Metadata\Get;
use ApiPlatform\Metadata\Post;
use App\Domain\Project\Entity\Project;
use App\Domain\User\Repository\UserRepository;
use App\Presentation\Controller\User\RegisterController;
use App\Shared\Domain\Trait\UuidTrait;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\Common\Collections\Collection;
use Doctrine\DBAL\Types\Types;
use Doctrine\ORM\Mapping as ORM;
use Exception;
use Gedmo\Timestampable\Traits\TimestampableEntity;
use Override;
use Symfony\Component\Security\Core\User\PasswordAuthenticatedUserInterface;
use Symfony\Component\Security\Core\User\UserInterface;
use Symfony\Component\Serializer\Attribute\Groups;
use Symfony\Component\Uid\Uuid;

#[ORM\Entity(repositoryClass: UserRepository::class)]
#[ORM\Table(name: '`user`')]
#[ORM\UniqueConstraint(name: 'UNIQ_IDENTIFIER_EMAIL', fields: ['email'])]
#[ApiResource(
    operations: [
        new Get(
            normalizationContext: ['groups' => ['user:read']],
            uriTemplate: '/me',
        ),
        new Post(
            uriTemplate: '/auth/register',
            controller: RegisterController::class,
        ),
    ]
)]
class User implements UserInterface, PasswordAuthenticatedUserInterface
{
    use UuidTrait;
    use TimestampableEntity;
    public const GROUP_USER_READ = ['user:read', 'plan:read', 'subscription:read'];

    #[ORM\Column(length: 180)]
    #[Groups(['user:read'])]
    private ?string $email = null;

    #[ORM\Column(type: Types::STRING, nullable: true)]
    #[Groups(['user:read'])]
    private ?string $firstname = null;

    #[ORM\Column(type: Types::STRING, nullable: true)]
    #[Groups(['user:read'])]
    private ?string $lastname = null;

    #[ORM\Column(type: Types::STRING, nullable: true)]
    #[Groups(['user:read'])]
    private ?string $picture = null;

    /**
     * @var list<string> The user roles
     */
    #[ORM\Column]
    #[Groups(['user:read'])]
    private array $roles = [];

    #[ORM\Column]
    private ?string $password = null;

    private ?string $plainPassword = null;

    /**
     * @var Collection<int, Project>
     */
    #[ORM\OneToMany(targetEntity: Project::class, mappedBy: 'user')]
    private Collection $projects;

    public function __construct()
    {
        $this->id = Uuid::v7();
        $this->projects = new ArrayCollection();
    }

    #[Groups(['user:read'])]
    public function getId(): Uuid
    {
        return $this->id;
    }

    public function getEmail(): string
    {
        if (null === $this->email) {
            throw new Exception('Email not found');
        }

        return $this->email;
    }

    public function getName(): string
    {
        return $this->firstname . ' ' . $this->lastname;
    }

    /**
     * @return non-empty-string
     *
     * @see UserInterface
     */
    #[Override]
    public function getUserIdentifier(): string
    {
        $identifier = (string) $this->email;
        if ('' === $identifier) {
            return 'unknown';
        }

        return $identifier;
    }

    /**
     * @see UserInterface
     */
    #[Override]
    public function getRoles(): array
    {
        $roles = $this->roles;
        $roles[] = 'ROLE_USER';

        return array_unique($roles);
    }

    /**
     * @see PasswordAuthenticatedUserInterface
     */
    #[Override]
    public function getPassword(): ?string
    {
        return $this->password;
    }

    public function getPlainPassword(): ?string
    {
        return $this->plainPassword;
    }

    /**
     * @see UserInterface
     */
    public function eraseCredentials(): void
    {
        $this->plainPassword = null;
    }

    public function getFirstname(): ?string
    {
        return $this->firstname;
    }

    public function getLastname(): ?string
    {
        return $this->lastname;
    }

    public function setFirstname(string $firstname): static
    {
        $this->firstname = $firstname;

        return $this;
    }

    public function setLastname(string $lastname): static
    {
        $this->lastname = $lastname;

        return $this;
    }

    public function setEmail(string $email): static
    {
        if (null !== $this->email && $this->email !== $email) {
            return $this;
        }

        $this->email = $email;

        return $this;
    }

    /**
     * @param list<string> $roles
     */
    public function setRoles(array $roles): static
    {
        $this->roles = $roles;

        return $this;
    }

    public function setPlainPassword(string $password): static
    {
        $this->plainPassword = $password;

        return $this;
    }

    public function setPassword(string $password): static
    {
        $this->password = $password;

        return $this;
    }

    public function getPicture(): ?string
    {
        return $this->picture;
    }

    public function setPicture(string $picture): static
    {
        $this->picture = $picture;

        return $this;
    }

    /**
     * @return Collection<int, Project>
     */
    public function getProjects(): Collection
    {
        return $this->projects;
    }

    /**
     * @param Collection<int, Project> $projects
     */
    public function setProjects(Collection $projects): void
    {
        $this->projects = $projects;
    }

    public function addProject(Project $project): void
    {
        $this->projects->add($project);
    }

    public function removeProject(Project $project): void
    {
        $this->projects->removeElement($project);
    }

    public function getProject(): Project
    {
        $project = $this->projects->filter(fn (Project $project): bool => $project->isActive())->first();

        if (false === $project) {
            throw new Exception('Project not found');
        }

        return $project;
    }
}
