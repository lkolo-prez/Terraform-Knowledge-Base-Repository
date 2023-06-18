Tytuł: Terraform: Zmienne i wyrażenia - elastyczność i łatwiejsze zarządzanie konfiguracją

Scenariusz filmu:

Scena 1: Wstęp
- Tło: biuro, stół z laptopem, telefonem i notatnikami.
- Narrator: "Terraform to potężne narzędzie do zarządzania infrastrukturą jako kod. W tym filmie dowiesz się, jak wykorzystać zmienne i wyrażenia w Terraform, aby uzyskać większą elastyczność i łatwiej zarządzać swoją konfiguracją."

Scena 2: Co to są zmienne w Terraform?
- Tło: biuro, stół z laptopem wyświetlającym kod Terraform.
- Narrator: "Zmienne to parametry, które pozwalają na modyfikowanie wartości w konfiguracji Terraform bez konieczności zmiany kodu. Zmienne sprawiają, że konfiguracja staje się bardziej elastyczna i łatwiejsza do zarządzania."
  - Pokaż przykład deklaracji zmiennej:
    ```
    variable "instance_type" {
      description = "Typ instancji EC2"
      default     = "t2.micro"
    }
    ```

Scena 3: Korzystanie ze zmiennych
- Tło: biuro, stół z laptopem wyświetlającym kod Terraform.
- Narrator: "Aby użyć zmiennej w konfiguracji Terraform, użyj składni 'var.nazwa_zmiennej'."
  - Pokaż przykład użycia zmiennej:
    ```
    resource "aws_instance" "example" {
      ami           = "ami-0c55b159cbfafe1f0"
      instance_type = var.instance_type

      tags = {
        Name = "example-instance"
      }
    }
    ```

Scena 4: Co to są wyrażenia w Terraform?
- Tło: biuro, stół z laptopem wyświetlającym kod Terraform.
- Narrator: "Wyrażenia to konstrukcje języka Terraform, które pozwalają na operowanie wartościami w konfiguracji. Przykłady wyrażeń to referencje do zmiennych, funkcje, czy też operacje matematyczne i logiczne."
  - Pokaż przykład wyrażenia:
    ```
    output "subnet_id" {
      value = aws_subnet.example.id
    }
    ```

Scena 5: Funkcje w Terraform
- Tło: biuro, stół z laptopem wyświetlającym kod Terraform.
- Narrator: "Terraform oferuje wiele funkcji, które umożliwiają przetwarzanie i manipulowanie danymi w konfiguracji. Przykładem funkcji jest 'length', która zwraca liczbę elementów w liście lub mapie."
  - Pokaż przykład użycia funkcji:
    ```
    output "vpc_cidr_block_length" {
      value = length(var.vpc_cidr_blocks)
    }
    ```

Scena 6: Podsumowanie
- Tło: biuro, stół z laptopem, telefonem i notatnikami.
- Narrator: "Zmienne i wyrażenia w Terraform pozwalają na tworzenie bardziej elastycznych, łatwiejszych do zarządzania i czytelnych konfiguracji infrastruktury jako kod. Dzięki nim, możemy tworzyć bardziej dynamiczne i skalowalne rozwiązania, dostosowując się do różnych scenariuszy."
  - Przypomnij o istotnych punktach:
    - "Zmienne pozwalają na parametryzację konfiguracji."
    - "Wyrażenia umożliwiają operacje na wartościach, takie jak funkcje, operacje matematyczne czy logiczne."
    - "Wykorzystanie zmiennych i wyrażeń sprawia, że kod Terraform jest czytelniejszy i łatwiejszy w utrzymaniu."

[Outro]
- Ekran z napisami "Dziękujemy za obejrzenie! Jeśli chcesz dowiedzieć się więcej o Terraform, subskrybuj nasz kanał i polub to wideo."