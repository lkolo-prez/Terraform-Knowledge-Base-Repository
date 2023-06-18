Tytuł: Terraform: Deklaratywny model programowania w zarządzaniu infrastrukturą

Scenariusz filmu:

Scena 1: Wstęp
- Tło: biuro, stół z laptopem, telefonem i notatnikami.
- Narrator: "Terraform to popularne narzędzie do zarządzania infrastrukturą jako kod. W tym filmie dowiesz się, jak wykorzystuje on deklaratywny model programowania do automatyzacji i ułatwienia zarządzania infrastrukturą."

Scena 2: Definicja deklaratywnego modelu programowania
- Tło: biuro, stół z laptopem wyświetlającym kod Terraform.
- Narrator: "W deklaratywnym modelu programowania opisujemy, jakiego rezultatu oczekujemy, a nie jak to osiągnąć. To podejście pozwala na skupienie się na celu, a nie na konkretnych krokach, co jest szczególnie przydatne w zarządzaniu infrastrukturą."

Scena 3: Przykład konfiguracji Terraform
- Tło: biuro, stół z laptopem wyświetlającym kod Terraform.
- Narrator: "Oto przykład deklaratywnej konfiguracji Terraform, która tworzy instancję serwera w chmurze AWS:"
  - Pokaż kod:
    ```
    provider "aws" {
      region = "us-west-2"
    }

    resource "aws_instance" "example" {
      ami           = "ami-0c55b159cbfafe1f0"
      instance_type = "t2.micro"

      tags = {
        Name = "example-instance"
      }
    }
    ```
- Narrator: "Jak widać, nie musimy podawać szczegółowych instrukcji, jak utworzyć instancję. Terraform sam zadba o to, aby nasza infrastruktura odpowiadała zdefiniowanemu stanowi."

Scena 4: Zalety deklaratywnego modelu programowania w Terraform
- Tło: biuro, stół z laptopem, telefonem i notatnikami.
- Narrator: "Deklaratywny model programowania ma wiele zalet:"
  - "Łatwiejsze zarządzanie zasobami: opisujemy tylko oczekiwany stan, a Terraform zajmuje się realizacją."
  - "Większa czytelność kodu: konfiguracje są bardziej zrozumiałe, co ułatwia współpracę zespołową."
  - "Mniejsza ilość błędów: zmniejsza się ryzyko wprowadzenia błędów przez użytkowników podczas ręcznego wykonywania poleceń."

Scena 5: Planowanie zmian i użycie komendy 'terraform apply'
- Tło: biuro, stół z laptopem z konsolą Terraform.
- Narrator: "Terraform pozwala na planowanie zmian w infrastrukturze przed ich wprowadzeniem. Wystarczy użyć komendy 'terraform plan', aby zobaczyć spodziewane zmiany w zasobach. Następnie, możemy zastosować te zmiany za pomocą komendy 'terraform apply'."
  - Pokaż komendy:
    ```
    terraform plan
    terraform apply
    ```

Scena 6: Deklaratywność w tworzeniu modułów
- Tło: biuro, stół z laptopem wyświetlającym kod Terraform z modułami.
- Narrator: "Deklaratywność w Terraform sprawia, że tworzenie modułów i ich konfiguracja jest prostsza. Moduły pozwalają na reużywanie kodu oraz utrzymanie czytelności konfiguracji."
  - Pokaż przykład modułu:
    ```
    module "vpc" {
      source = "terraform-aws-modules/vpc/aws"

      name = "my-vpc"
      cidr = "10.0.0.0/16"

      azs             = ["us-west-2a", "us-west-2b"]
      private_subnets = ["10.0.1.0/24", "10.0.2.0/24"]
      public_subnets  = ["10.0.101.0/24", "10.0.102.0/24"]

      enable_nat_gateway = true
      enable_vpn_gateway = true
    }
    ```

Scena 7: Idempotentność w Terraform
- Tło: biuro, stół z laptopem, telefonem i notatnikami.
- Narrator: "Deklaratywny model programowania sprawia, że Terraform jest idempotentny. Oznacza to, że niezależnie od tego, ile razy uruchomimy komendę 'terraform apply', otrzymamy ten sam oczekiwany stan infrastruktury. Dzięki temu unikamy nieoczekiwanych zmian i problemów w naszej infrastrukturze."

Scena 8: Podsumowanie
- Tło: biuro, stół z laptopem, telefonem i notatnikami.
- Narrator: "Terraform to potężne narzędzie do zarządzania infrastrukturą jako kod, które wykorzystuje deklaratywny model programowania. Dzięki temu mamy pełną kontrolę nad stanem naszej infrastruktury, przy jednoczesnym uproszczeniu procesu zarządzania nią. Terraform sprawia, że zarządzanie infrastrukturą staje się prostsze, bardziej elastyczne i mniej podatne na błędy."

[Outro]
- Ekran z napisami "Dziękujemy za obejrzenie! Jeśli chcesz dowiedzieć się więcej o Terraform, subskrybuj nasz kanał i polub to wideo."

