# Dokumentacja

## 1. Wprowadzenie

Krótki opis, struktura projektu, oraz wymagania systemowe są przedstawione [tutaj](README.md).

## 2. Kluczowe algorytmy i mechanizmy

### 2.1. System energetyczny
- Każde zwierzę posiada poziom energii, który maleje z czasem
- Pasywne zużycie energii rośnie z wiekiem zwierzęcia
- Króliki odzyskują energię jedząc trawę
- Lisy odzyskują energię polując na króliki
- System energetyczny tworzy naturalną presję na zwierzęta do poszukiwania pożywienia

### 2.2. Mechanizm rozmnażania
- Zwierzęta mogą się rozmnażać tylko gdy:
    - Posiadają wystarczający poziom energii
    - Minął okres regeneracji po poprzednim rozmnożeniu
    - Znajdują się w bezpośrednim sąsiedztwie innego osobnika tego samego gatunku
- Okres regeneracji zapobiega nadmiernej eksplozji populacji

### 2.3. System ruchu
- Zwierzęta poruszają się losowo w obrębie sąsiednich pól
- Implementacja uwzględnia kolizje - zwierzęta nie mogą zajmować tego samego pola
- System granic mapy zapewnia, że zwierzęta pozostają w obszarze symulacji

### 2.4. Wzrost i rozprzestrzenianie się trawy
- Trawa rośnie stopniowo do maksymalnej wartości
- Trawa rozprzestrzenia się na sąsiednie pola z określonym prawdopodobieństwem
- Mechanizm zapewnia regenerację zasobów pokarmowych dla królików

## 3. Obserwacje i wnioski

### 3.1. Dynamika populacji

- Zaobserwowano cykliczne wahania liczebności populacji
- Typowy cykl:
  1. Wzrost populacji królików przy dostępności trawy
  2. Wzrost populacji lisów w odpowiedzi na dużą liczbę królików
  3. Spadek populacji królików przez drapieżnictwo
  4. Spadek populacji lisów z powodu braku pożywienia
- Najczęssze przypadki końcowe:
  - Lisy zjadają wszystkie króliki i umierają z głodu
  - Zostaje jeden królik który biega dopóki nie zdechnie ze starości

### 3.2. Ciekawe zjawiska
- Tworzenie się "hot spotów" - obszarów o zwiększonej aktywności
- Okresowe wymieranie populacji przy skrajnych parametrach
- Samoregulacja systemu przy odpowiednich wartościach parametrów

### 3.3. Punkty krytyczne
- Zbyt mała początkowa populacja królików może prowadzić do wymarcia lisów
- Zbyt duża populacja lisów może doprowadzić do załamania całego ekosystemu

## 4. Ocena i wnioski

### 4.1. Mocne strony symulacji
- Realistyczne odwzorowanie podstawowych mechanizmów ekosystemu
- Dane i wykres umożliwiają łatwe prowadzenie obserwacji

### 4.2. Potencjalne rozszerzenia
- Implementacja bardziej zaawansowanych zachowań zwierząt (np. śledzenie ofiar przez drapieżnika)

### 4.3 Podsumowanie
Symulacja stanowi ciekawy przykład prostego modelu ekosystemu, który pozwala na obserwację interakcji między gatunkami i dynamiki populacji. Mimo swojej prostoty, oferuje wiele możliwości do dalszego rozwoju i badań nad bardziej rozbudowanymi systemami. Zadanie było wymagające, ale satysfakcjonujące. Nawet jeśli nie udało mi się znaleźć parametrów, które zapewniłyby balans i przetrwanie obu gatunków, zdarzały się kilkukrotne powtórzenia cyklu opisanego w **3.1**.
